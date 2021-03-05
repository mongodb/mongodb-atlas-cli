// Copyright 2021 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package quickstart

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongocli/internal/convert"
	"github.com/mongodb/mongocli/internal/net"
	"github.com/mongodb/mongocli/internal/usage"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

func newAccessListQuestion(publicIP string) *survey.Question {
	if publicIP != "" {
		return nil
	}

	message := "Insert the IP entry to add to the Access List"
	publicIP = net.IPAddress()

	if publicIP != "" {
		message = fmt.Sprintf("Insert the IP entry to add to the Access List [Press Enter to use your public IP address '%s']", publicIP)
	}

	return &survey.Question{
		Name: "ipAddress",
		Prompt: &survey.Input{
			Message: message,
			Help:    usage.AccessListIPEntry,
			Default: publicIP,
		},
	}
}

func newRegionQuestions(region, provider string) *survey.Question {
	if region != "" {
		return nil
	}
	return &survey.Question{
		Name: "region",
		Prompt: &survey.Select{
			Message: "Insert the physical location of your MongoDB cluster",
			Help:    usage.Region,
			Options: DefaultRegions[strings.ToUpper(provider)],
		},
	}
}

func (opts *Opts) newDBUserQuestions() []*survey.Question {
	var qs []*survey.Question
	if opts.DBUsername == "" {
		message := "Insert the Username for authenticating to MongoDB"
		usrDefault := dbUsername()
		if usrDefault != "" {
			message = fmt.Sprintf("Insert the Username for authenticating to MongoDB [Press Enter to use '%s']", usrDefault)
		}
		q := &survey.Question{
			Validate: func(val interface{}) error {
				username, _ := val.(string)
				user, err := opts.store.DatabaseUser(convert.AdminDB, opts.ConfigProjectID(), username)
				if err != nil {
					if !strings.Contains(err.Error(), fmt.Sprintf("No user with username %s exists.", username)) {
						return err
					}
				}

				if user != nil {
					return errors.New("a user with this username already exists")
				}

				return nil
			},
			Name: "dbUsername",
			Prompt: &survey.Input{
				Message: message,
				Help:    usage.DBUsername,
				Default: usrDefault,
			},
		}

		qs = append(qs, q)
	}

	if opts.DBUserPassword == "" {
		q := &survey.Question{
			Name: "dbUserPassword",
			Prompt: &survey.Password{
				Message: "Insert the Password for authenticating to MongoDB [Press Enter to use an auto-generated password]",
				Help:    usage.Password,
			},
		}

		qs = append(qs, q)
	}

	return qs
}

func (opts *Opts) newClusterQuestions() []*survey.Question {
	var qs []*survey.Question

	if opts.ClusterName == "" {
		clusterName := opts.newClusterName()
		message := "Insert the cluster name"
		if clusterName != "" {
			message = fmt.Sprintf("Insert the cluster name [Press Enter to use the auto-generated name '%s']", clusterName)
		}
		q := &survey.Question{
			Name: "clusterName",
			Prompt: &survey.Input{
				Message: message,
				Help:    usage.ClusterName,
				Default: clusterName,
			},
		}
		qs = append(qs, q)
	}

	if opts.Provider == "" {
		q := &survey.Question{
			Name: "provider",
			Prompt: &survey.Select{
				Message: "Insert the cloud service provider on which Atlas provisions the hosts",
				Help:    usage.Provider,
				Options: []string{"AWS", "GCP", "AZURE"},
			},
		}
		qs = append(qs, q)
	}

	return qs
}

// dbUsername returns the username of the user by running the command 'whoami'
func dbUsername() string {
	command := "whoami"
	cmd := exec.Command("bash", "-c", command)

	if runtime.GOOS == "windows" {
		cmd = exec.Command("powershell", "", command)
	}

	stdout, err := cmd.Output()

	if err != nil {
		return ""
	}

	// dbUsername can only contain ASCII letters, numbers, hyphens and underscores
	out := strings.TrimSpace(string(stdout))
	var re = regexp.MustCompile("([^A-Za-z0-9_-])")
	return re.ReplaceAllString(out, "_")
}

// newClusterName returns an auto-generate Cluster name
func (opts *Opts) newClusterName() string {
	cs, _ := opts.store.ProjectClusters(opts.ConfigProjectID(), nil)
	i := 0
	if clusters, ok := cs.([]atlas.Cluster); ok {
		for {
			clusterName := "QuickstartCluster" + strconv.Itoa(i)
			found := false
			for i := range clusters {
				if clusters[i].Name == clusterName {
					found = true
					break
				}
			}

			if !found {
				return clusterName
			}
			i++
		}
	}

	return ""
}
