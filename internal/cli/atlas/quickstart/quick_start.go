// Copyright 2020 MongoDB Inc
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
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongocli/e2e"
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/convert"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

const quickstartTemplate = "Now you can connect to your Atlas cluster with: mongo -u %s -p %s %s\n"

const (
	replicaSet        = "REPLICASET"
	diskSizeGB        = 10
	mdbVersion        = "4.2"
	shards            = 1
	tier              = "M10"
	members           = 3
	zoneName          = "Zone 1"
	accessListComment = "IP added with mongocli atlas quickstart"
	atlasAdmin        = "atlasAdmin"
	none              = "NONE"
	passwordLength    = 12
	maxRandNum        = 10000
)

type Opts struct {
	cli.GlobalOpts
	cli.WatchOpts
	clusterName      string
	provider         string
	region           string
	ipAddress        string
	dbUsername       string
	dbUserPassword   string
	connectionString string
	store            store.AtlasClusterQuickStarter
}

func (opts *Opts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *Opts) Run() error {
	if _, err := opts.store.CreateCluster(opts.newCluster()); err != nil {
		return err
	}

	// Add IP to project’s IP access list
	entry := opts.newWhitelist()
	if _, err := opts.store.CreateProjectIPAccessList(entry); err != nil {
		return err
	}

	// Create DBUser
	if er := opts.createDatabaseUser(); er != nil {
		return er
	}

	fmt.Println("Creating your cluster...")
	if er := opts.Watch(opts.watcher); er != nil {
		return er
	}

	// Get cluster's connection string
	cluster, err := opts.store.AtlasCluster(opts.ConfigProjectID(), opts.clusterName)
	if err != nil {
		return err
	}
	opts.connectionString = cluster.SrvAddress

	fmt.Printf(quickstartTemplate, opts.dbUsername, opts.dbUserPassword, opts.connectionString)
	return nil
}

func (opts *Opts) watcher() (bool, error) {
	result, err := opts.store.AtlasCluster(opts.ConfigProjectID(), opts.clusterName)
	if err != nil {
		return false, err
	}
	return result.StateName == "IDLE", nil
}

func (opts *Opts) createDatabaseUser() error {
	user, err := opts.store.DatabaseUser(convert.AdminDB, opts.ConfigProjectID(), opts.dbUsername)
	if err != nil {
		if !strings.Contains(err.Error(), fmt.Sprintf("No user with username %s exists.", opts.dbUsername)) {
			return err
		}
	}

	if user != nil {
		return nil
	}

	// Create dbUser
	if _, err := opts.store.CreateDatabaseUser(opts.newDatabaseUser()); err != nil {
		return err
	}

	return nil
}

func (opts *Opts) newDatabaseUser() *atlas.DatabaseUser {
	return &atlas.DatabaseUser{
		Roles:        convert.BuildAtlasRoles([]string{atlasAdmin}),
		GroupID:      opts.ConfigProjectID(),
		Password:     opts.dbUserPassword,
		X509Type:     none,
		AWSIAMType:   none,
		LDAPAuthType: none,
		DatabaseName: convert.AdminDB,
		Username:     opts.dbUsername,
	}
}

// newIPAddress returns client's public ip
func newIPAddress() (string, error) {
	publicIP := ""
	for _, uri := range APIURIs {
		req, err := http.NewRequestWithContext(
			context.Background(),
			http.MethodGet,
			uri,
			nil,
		)

		req.Header.Add("Accept", "application/json")

		if err == nil {
			res, err := http.DefaultClient.Do(req)

			if err == nil {
				responseBytes, err := ioutil.ReadAll(res.Body)
				res.Body.Close()
				if err == nil {
					publicIP = string(responseBytes)
					break
				}
			}
		}
	}

	if publicIP == "" {
		return publicIP, errors.New("error in finding your public IP, please use --ip to provide your public ip")
	}

	return publicIP, nil
}

// APIURIs is the URIs of the services used by newIPAddress to get the client's public IP.
var APIURIs = []string{
	"https://api.ipify.org",
	"http://myexternalip.com/raw",
	"http://ipinfo.io/ip",
	"http://ipecho.net/plain",
	"http://icanhazip.com",
	"http://ifconfig.me/ip",
	"http://ident.me",
	"http://checkip.amazonaws.com",
	"http://bot.whatismyipaddress.com",
	"http://whatismyip.akamai.com",
	"http://wgetip.com",
	"http://ip.tyk.nu",
}

func (opts *Opts) newWhitelist() *atlas.ProjectIPWhitelist {
	return &atlas.ProjectIPWhitelist{
		GroupID:   opts.ConfigProjectID(),
		Comment:   accessListComment,
		IPAddress: opts.ipAddress,
	}
}

func (opts *Opts) newCluster() *atlas.Cluster {
	diskSizeGB := float64(diskSizeGB)
	return &atlas.Cluster{
		GroupID:             opts.ConfigProjectID(),
		ClusterType:         replicaSet,
		ReplicationSpecs:    []atlas.ReplicationSpec{opts.newReplicationSpec()},
		ProviderSettings:    opts.newProviderSettings(),
		MongoDBMajorVersion: mdbVersion,
		DiskSizeGB:          &diskSizeGB,
		Name:                opts.clusterName,
	}
}

func (opts *Opts) newReplicationSpec() atlas.ReplicationSpec {
	var (
		readOnlyNodes int64 = 0
		priority      int64 = 7
		shards        int64 = shards
		members       int64 = members
	)
	replicationSpec := atlas.ReplicationSpec{
		NumShards: &shards,
		ZoneName:  zoneName,
		RegionsConfig: map[string]atlas.RegionsConfig{
			opts.region: {
				ReadOnlyNodes:  &readOnlyNodes,
				ElectableNodes: &members,
				Priority:       &priority,
			},
		},
	}
	return replicationSpec
}

func (opts *Opts) newProviderSettings() *atlas.ProviderSettings {
	return &atlas.ProviderSettings{
		InstanceSizeName: tier,
		ProviderName:     opts.provider,
		RegionName:       opts.region,
	}
}

// askRequiredFlags allows the user to set required flags by using interactive prompts
func (opts *Opts) askRequiredFlags() error {
	if opts.dbUsername == "" {
		dbUsername := ""
		dbUsernamePrompt := &survey.Input{
			Message: "Insert the Username for authenticating to MongoDB [Press Enter to use an autogenerated username]",
			Help:    usage.DBUsername,
		}

		if err := survey.AskOne(dbUsernamePrompt, &dbUsername); err != nil {
			return err
		}
		opts.dbUsername = dbUsername

		if opts.dbUsername == "" {
			n, err := e2e.RandInt(maxRandNum)
			if err != nil {
				return err
			}
			opts.dbUsername = "quickStart_" + n.String()
		}
	}

	if opts.dbUserPassword == "" {
		dbPassword := ""
		dbPasswordPrompt := &survey.Password{
			Message: "Insert the Password for authenticating to MongoDB [Press Enter to use an autogenerated password]",
			Help:    usage.Password,
		}

		if err := survey.AskOne(dbPasswordPrompt, &dbPassword); err != nil {
			return err
		}

		opts.dbUserPassword = dbPassword
		if opts.dbUserPassword == "" {
			p, err := newAutogeneratedPassword(passwordLength)
			if err != nil {
				return err
			}
			opts.dbUserPassword = p
		}
	}

	if opts.ipAddress == "" {
		answer := ""
		publicIP, err := newIPAddress()
		message := "Insert the IP entry to add to the Access List"
		if err == nil {
			message = fmt.Sprintf(`Insert the IP entry to add to the Access List [Press Enter to use your public IP "%s"]`, publicIP)
		}

		publicIPPrompt := survey.Input{
			Message: message,
			Help:    usage.AccessListIPEntry,
		}

		if err := survey.AskOne(&publicIPPrompt, &answer); err != nil {
			return err
		}

		opts.ipAddress = answer
		if opts.ipAddress == "" {
			opts.ipAddress = publicIP
		}
	}

	err := opts.askProviderAndRegionFlags()
	if err != nil {
		return err
	}

	return nil
}

func (opts *Opts) askProviderAndRegionFlags() error {
	if opts.provider == "" {
		providerPrompt := &survey.Select{
			Message: "Insert the cloud service provider on which Atlas provisions the hosts",
			Help:    usage.Provider,
			Options: []string{"AWS", "GCP", "AZURE"},
		}

		if err := survey.AskOne(providerPrompt, &opts.provider); err != nil {
			return err
		}
	}

	if opts.region == "" {
		regionOption := []string{"US_EAST_1", "US_WEST_2", "AP_SOUTH_1", "AP_EAST_2", "EU_WEST_1", "EU_CENTRAL_1", "ME_SOUTH_1", "AF_SOUTH_1"}
		if opts.provider == "AZURE" {
			regionOption = []string{"US_EAST_2", "US_WEST", "EUROPE_NORTH"}
		}

		if opts.provider == "GCP" {
			regionOption = []string{"CENTRAL_US", "CANADA_CENTRAL", "WESTERN_EUROPE", "ASIA_SOUTH_EAST", "SOUTH_AFRICA_NORTH", "UAE_NORTH"}
		}

		regionPrompt := &survey.Select{
			Message: "Insert the physical location of your MongoDB cluster",
			Help:    usage.Region,
			Options: regionOption,
		}

		if err := survey.AskOne(regionPrompt, &opts.region); err != nil {
			return err
		}
	}
	return nil
}

func newAutogeneratedPassword(length int) (string, error) {
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	var b strings.Builder
	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", err
		}
		b.WriteRune(chars[index.Int64()])
	}

	return b.String(), nil
}

// mongocli atlas dbuser(s) quickstart [--clusterName clusterName] [--provider provider] [--region regionName] [--projectId projectId] [--username username] [--password password]
func Builder() *cobra.Command {
	opts := &Opts{}
	cmd := &cobra.Command{
		Use: "quickstart",
		Example: `mongocli atlas quickstart
mongocli atlas quickstart --clusterName Test --provider GPC --username dbuserTest --password Test!
`,
		Short: QuickStart,
		Long:  LongQuickStart,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.askRequiredFlags,
				opts.initStore,
				opts.InitOutput(cmd.OutOrStdout(), ""),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "GetStarted", usage.ClusterName)
	cmd.Flags().StringVar(&opts.provider, flag.Provider, "", usage.Provider)
	cmd.Flags().StringVarP(&opts.region, flag.Region, flag.RegionShort, "", usage.Region)
	cmd.Flags().StringVar(&opts.ipAddress, flag.IP, "", usage.AccessListIPEntry)
	cmd.Flags().StringVar(&opts.dbUsername, flag.Username, "", usage.DBUsername)
	cmd.Flags().StringVar(&opts.dbUserPassword, flag.Password, "", usage.Password)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}
