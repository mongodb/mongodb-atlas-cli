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
	"github.com/10gen/mcli/internal/flags"
	"github.com/10gen/mcli/internal/usage"
	"github.com/10gen/mcli/internal/utils"
	"github.com/spf13/cobra"
)

type configSetOpts struct {
	*globalOpts
	prop string
	val  string
}

func (opts *configSetOpts) Run() error {
	opts.Config.Set(opts.prop, opts.val)
	if err := opts.Config.Save(); err != nil {
		return err
	}
	fmt.Printf("Updated prop '%s'\n", opts.prop)
	return nil
}

func ConfigSetBuilder() *cobra.Command {
	opts := &configSetOpts{
		globalOpts: newGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:   "set [prop] [val]",
		Short: "Configure the tool.",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return fmt.Errorf("accepts %d arg(s), received %d", 2, len(args))
			}
			if !utils.StringInSlice(cmd.ValidArgs, args[0]) {
				return fmt.Errorf("invalid prop %q", args[0])
			}
			return nil
		},
		ValidArgs: config.Properties(),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.globalOpts.loadConfig()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.prop = args[0]
			opts.val = args[1]
			return opts.Run()
		},
	}
	cmd.Flags().StringVarP(&opts.profile, flags.Profile, flags.ProfileShort, config.DefaultProfile, usage.ProfileConfig)

	return cmd
}
