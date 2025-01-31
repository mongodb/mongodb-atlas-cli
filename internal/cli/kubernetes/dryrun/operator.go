// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dryrun

import (
	"context"
	"fmt"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/resources"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"time"
)

type Worker struct {
	targetNamespace string
	watchNamespaces string
	wait            bool
	akoVersion      string
	waitSec         int64
}

func NewWorker() *Worker {
	return &Worker{}
}

func (r *Worker) WithTargetNamespace(targetNamespace string) *Worker {
	r.targetNamespace = targetNamespace
	return r
}

func (r *Worker) WithWatchNamespaces(watchNamespaces string) *Worker {
	r.watchNamespaces = watchNamespaces
	return r
}

func (r *Worker) WithOperatorVersion(operatorVersion string) *Worker {
	r.akoVersion = operatorVersion
	return r
}

func (r *Worker) WithWaitForCompletion(waitForCompletion bool) *Worker {
	r.wait = waitForCompletion
	return r
}

func (r *Worker) WithWaitTimeoutSec(waitSec int64) *Worker {
	r.waitSec = waitSec
	return r
}

func (r *Worker) Run() error {
	conf, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("failed to get k8s config: %w", err)
	}

	c, err := client.New(conf, client.Options{})
	if err != nil {
		return fmt.Errorf("failed to create client: %w", err)
	}

	jb := &batchv1.Job{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Job",
			APIVersion: "batch/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: resources.NormalizeAtlasName("ako-dry-run-", resources.AtlasNameToKubernetesName()),
			Namespace:    r.targetNamespace,
		},
		Spec: batchv1.JobSpec{
			BackoffLimit: makePtr[int32](1),
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					ServiceAccountName: "mongodb-atlas-operator",
					RestartPolicy:      corev1.RestartPolicyNever,
					Containers: []corev1.Container{
						{
							Name: "ako-dry-run",
							//Image: fmt.Sprintf("quay.io/mongodb/atlas-kubernetes-operator:%s", r.akoVersion),
							Image:   "docker.io/ikarpukhin/mongodb-atlas-kubernetes:dry-run",
							Command: []string{"/manager"},
							Args: []string{
								"--atlas-domain=https://cloud-qa.mongodb.com/",
								"--log-level=info",
								"--log-encoder=json",
								"--dry-run",
							},
							Env: []corev1.EnvVar{
								{
									Name:  "OPERATOR_POD_NAME",
									Value: "ako-dry-run",
								},
								{
									Name:  "OPERATOR_NAMESPACE",
									Value: r.targetNamespace,
								},
								{
									Name:  "WATCH_NAMESPACE",
									Value: r.targetNamespace,
								},
								{
									Name:  "JOB_NAME",
									Value: "ako-dry-run",
								},
							},
							LivenessProbe: &corev1.Probe{
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/healthz",
										Port: intstr.IntOrString{IntVal: 8081},
									},
								},
								InitialDelaySeconds: 15,
								PeriodSeconds:       20,
							},
							ImagePullPolicy: "Always",
						},
					},
				},
			},
		},
	}

	if err := c.Create(context.Background(), jb); err != nil {
		return fmt.Errorf("failed to create job: %w", err)
	}

	fmt.Printf("AKO dry run job '%s' created successfully at '%s'\r\n",
		jb.Name, jb.CreationTimestamp.Format(time.DateTime))

	if !r.wait {
		return nil
	}

	ctx, timeoutF := context.WithTimeout(context.Background(), time.Duration(r.waitSec)*time.Second)
	defer timeoutF()

	if err := waitForJob(ctx, c, jb); err != nil {
		return fmt.Errorf("failed to wait for job: %w", err)
	}

	fmt.Printf("AKO dry run job '%s' completed successfully at '%s'\r\n",
		jb.Name, time.Now().Format(time.DateTime))
	return nil
}

func waitForJob(ctx context.Context, c client.Client, job *batchv1.Job) error {
	attempts := 0
	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("timeout: job did not complete within the expected time: %w", ctx.Err())
		default:
			jb := &batchv1.Job{}
			if err := c.Get(ctx, client.ObjectKey{Name: job.Name, Namespace: job.Namespace}, jb); err != nil {
				return fmt.Errorf("failed to get job: %w", err)
			}

			if jb.Status.Succeeded > 0 {
				return nil
			}

			if jb.Status.Failed > 0 {
				return fmt.Errorf("job failed with conditions: %+v", jb.Status.Conditions)
			}

			time.Sleep(5 * time.Second)
			attempts++
			fmt.Printf("Waiting for job to complete... Attempt #%d\r\n", attempts)
		}
	}
}

func makePtr[T any](v T) *T {
	return &v
}
