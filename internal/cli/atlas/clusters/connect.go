// Copyright 2023 MongoDB Inc
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

package clusters

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/denisbrodbeck/machineid"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/mongosh"
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20230201008/admin"
)

type ConnectOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	name  string
	store store.Connect
	fs    afero.Fs
}

func (opts *ConnectOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func username() (string, error) {
	id, err := machineid.ProtectedID("atlascli")
	if err != nil {
		return "", err
	}
	username := "atlascli_" + id

	return username, nil
}

func (opts *ConnectOpts) invalidCertFile(certPath string) bool {
	aMonthAgo := time.Now().AddDate(0, -1, 0)
	info, err := opts.fs.Stat(certPath)

	return os.IsNotExist(err) || info.ModTime().Before(aMonthAgo)
}

func (opts *ConnectOpts) userExists(usr string) bool {
	_, err := opts.store.DatabaseUser("$external", opts.ConfigProjectID(), usr)
	return err == nil
}

func (opts *ConnectOpts) createUser(usr string) error {
	_, err := opts.store.CreateDatabaseUser(&admin.CloudDatabaseUser{
		Username:     usr,
		DatabaseName: "$external",
		GroupId:      opts.ConfigProjectID(),
		X509Type:     pointer.Get("MANAGED"),
		Roles:        []admin.DatabaseUserRole{{RoleName: "readWriteAnyDatabase", DatabaseName: "admin"}},
	})

	return err
}

func (opts *ConnectOpts) certPath() (string, error) {
	configDir, err := config.AtlasCLIConfigHome()
	if err != nil {
		return "", err
	}

	certPath := path.Join(configDir, opts.ConfigProjectID()+".pem")
	return certPath, nil
}

func (opts *ConnectOpts) connectionString(certPath string) (string, error) {
	r, err := opts.store.AtlasCluster(opts.ConfigProjectID(), opts.name)
	if err != nil {
		return "", err
	}

	if r == nil || r.ConnectionStrings == nil || r.ConnectionStrings.StandardSrv == nil {
		return "", errors.New("connection string not found")
	}

	cnnStr := *r.ConnectionStrings.StandardSrv
	if strings.Contains(cnnStr, "?") {
		cnnStr += "&"
	} else {
		cnnStr += "?"
	}
	cnnStr += "authSource=$external&authMechanism=MONGODB-X509&tls=true&tlsCertificateKeyFile=" + certPath

	return cnnStr, nil
}

func (opts *ConnectOpts) Run(ctx context.Context) error {
	certPath, err := opts.certPath()
	if err != nil {
		return err
	}

	cnnStr, err := opts.connectionString(certPath)
	if err != nil {
		return err
	}

	if opts.invalidCertFile(certPath) {
		usr, err := username()
		if err != nil {
			return err
		}

		if !opts.userExists(usr) {
			if err := opts.createUser(usr); err != nil {
				return err
			}
		}

		cert, err := opts.store.CreateDBUserCertificate(opts.ConfigProjectID(), usr, 1)
		if err != nil {
			return err
		}

		if err := afero.WriteFile(opts.fs, certPath, []byte(cert), os.ModePerm); err != nil {
			return err
		}
	}

	return mongosh.Run("", "", cnnStr)
}

// atlas clusters connect [clusterName].
func ConnectBuilder() *cobra.Command {
	opts := &ConnectOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "connect <clusterName>",
		Short: "Connect to the specified cluster.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"clusterNameDesc": "Name of the cluster to connect.",
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), "{{.StandardSrv}}"),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.name = args[0]
			return opts.Run(cmd.Context())
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
