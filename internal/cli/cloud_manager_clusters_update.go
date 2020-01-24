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

	"github.com/10gen/mcli/internal/config"
	"github.com/10gen/mcli/internal/convert"
	"github.com/10gen/mcli/internal/flags"
	"github.com/10gen/mcli/internal/messages"
	"github.com/10gen/mcli/internal/search"
	"github.com/10gen/mcli/internal/store"
	"github.com/10gen/mcli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

type cmClustersUpdateOpts struct {
	*globalOpts
	filename string
	fs       afero.Fs
	store    store.AutomationStore
}

func (opts *cmClustersUpdateOpts) init() error {
	if opts.ProjectID() == "" {
		return errMissingProjectID
	}

	s, err := store.New()

	if err != nil {
		return err
	}

	opts.store = s
	return nil
}

func (opts *cmClustersUpdateOpts) Run() error {
	newConfig, err := convert.NewClusterConfigFromFile(opts.fs, opts.filename)
	if err != nil {
		return err
	}
	current, err := opts.store.GetAutomationConfig(opts.ProjectID())

	if err != nil {
		return err
	}

	if !search.ClusterExists(current, newConfig.Name) {
		return fmt.Errorf("cluster '%s' doesn't exist", newConfig.Name)
	}

	err = newConfig.PatchAutomationConfig(current)

	if err != nil {
		return err
	}

	if err = opts.store.UpdateAutomationConfig(opts.ProjectID(), current); err != nil {
		return err
	}

	fmt.Print(messages.DeploymentStatus(config.OpsManagerURL(), opts.ProjectID()))

	return nil
}

// mcli cloud-manager cluster(s) update --projectId projectId --file myfile.yaml
func CloudManagerClustersUpdateBuilder() *cobra.Command {
	opts := &cmClustersUpdateOpts{
		globalOpts: newGlobalOpts(),
		fs:         afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a Cloud Manager cluster.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVarP(&opts.filename, flags.File, flags.FileShort, "", "Filename to use to update the cluster")

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	_ = cmd.MarkFlagRequired(flags.File)

	return cmd
}
