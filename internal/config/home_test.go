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

package config

import (
	"os"
	"testing"

	"github.com/mitchellh/go-homedir"
)

func TestConfig_configHome(t *testing.T) {
	t.Run("with XDG_CONFIG_HOME", func(t *testing.T) {
		xdgHome := "my_config"
		_ = os.Setenv("XDG_CONFIG_HOME", xdgHome)
		home, err := configHome()
		if err != nil {
			t.Fatalf("configHome() unexpected error: %v", err)
		}
		if home != "my_config" {
			t.Errorf("configHome() = %s; want '%s'", home, xdgHome)
		}
		_ = os.Unsetenv("XDG_CONFIG_HOME")
	})
	t.Run("without XDG_CONFIG_HOME", func(t *testing.T) {
		homedir.DisableCache = true
		_ = os.Setenv("HOME", ".")
		home, err := configHome()
		if err != nil {
			t.Fatalf("configHome() unexpected error: %v", err)
		}
		if home != "./.config" {
			t.Errorf("configHome() = %s; want './.config'", home)
		}
		homedir.DisableCache = false
	})
}
