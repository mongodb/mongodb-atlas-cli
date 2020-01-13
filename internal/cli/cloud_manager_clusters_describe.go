// Copyright (C) 2020 - present MongoDB, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the Server Side Public License, version 1,
// as published by MongoDB, Inc.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// Server Side Public License for more details.
//
// You should have received a copy of the Server Side Public License
// along with this program. If not, see
// http://www.mongodb.com/licensing/server-side-public-license
//
// As a special exception, the copyright holders give permission to link the
// code of portions of this program with the OpenSSL library under certain
// conditions as described in each individual source file and distribute
// linked combinations including the program with the OpenSSL library. You
// must comply with the Server Side Public License in all respects for
// all of the code used other than as permitted herein. If you modify file(s)
// with this exception, you may extend this exception to your version of the
// file(s), but you are not obligated to do so. If you do not wish to do so,
// delete this exception statement from your version. If you delete this
// exception statement from all source files in the program, then also delete
// it in the license file.

package cli

import (
	"fmt"

	"github.com/10gen/mcli/internal/usage"

	"github.com/10gen/mcli/internal/config"
	"github.com/10gen/mcli/internal/convert"
	"github.com/10gen/mcli/internal/flags"
	"github.com/10gen/mcli/internal/store"
	"github.com/10gen/mcli/internal/utils"
	"github.com/spf13/cobra"
)

type cmClustersDescribeOpts struct {
	*globalOpts
	name  string
	store store.AutomationGetter
}

func (opts *cmClustersDescribeOpts) init() error {
	if err := opts.loadConfig(); err != nil {
		return err
	}

	if opts.ProjectID() == "" {
		return errMissingProjectID
	}

	s, err := store.New(opts.Config)

	if err != nil {
		return err
	}

	opts.store = s
	return nil
}

func (opts *cmClustersDescribeOpts) Run() error {
	result, err := opts.store.GetAutomationConfig(opts.ProjectID())

	if err != nil {
		return err
	}

	clusterConfigs := convert.FromAutomationConfig(result)
	for _, rs := range clusterConfigs {
		if rs.Name == opts.name {
			return utils.PrettyJSON(rs)
		}

	}
	return fmt.Errorf("replicaset %s not found", opts.name)
}

// mcli cloud-manager cluster(s) describe [name] --projectId projectId
func CloudManagerClustersDescribeBuilder() *cobra.Command {
	opts := &cmClustersDescribeOpts{
		globalOpts: newGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:   "describe [name]",
		Short: "Command to describe a Cloud Manager cluster.",
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.name = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVarP(&opts.profile, flags.Profile, "p", config.DefaultProfile, usage.Profile)

	return cmd
}
