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
	"fmt"

	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

type Profile struct {
	Name      string
	configDir string
	fs        afero.Fs
}

func New(name string) (*Profile, error) {
	configDir, err := configHome()
	if err != nil {
		return nil, err
	}

	p := new(Profile)
	p.Name = name
	p.configDir = configDir
	p.fs = afero.NewOsFs()

	return p, nil
}

func Properties() []string {
	return []string{service, publicAPIKey, privateAPIKey, opsManagerURL, baseURL}
}

func (p Profile) Set(name, value string) {
	viper.Set(fmt.Sprintf("%s.%s", p.Name, name), value)
}

func (p Profile) GetString(name string) string {
	if viper.IsSet(name) && viper.GetString(name) != "" {
		return viper.GetString(name)
	}

	return viper.GetString(fmt.Sprintf("%s.%s", p.Name, name))
}

// Service get configured service
func (p Profile) Service() string {
	if viper.IsSet(service) {
		return viper.GetString(service)
	}
	serviceKey := fmt.Sprintf("%s.%s", p.Name, service)
	if viper.IsSet(serviceKey) {
		return viper.GetString(serviceKey)
	}
	return CloudService
}

// SetService set configured service
func (p Profile) SetService(v string) {
	p.Set(service, v)
}

// PublicAPIKey get configured public api key
func (p Profile) PublicAPIKey() string {
	return p.GetString(publicAPIKey)
}

// SetPublicAPIKey set configured publicAPIKey
func (p Profile) SetPublicAPIKey(v string) {
	p.Set(publicAPIKey, v)
}

// PrivateAPIKey get configured private api key
func (p Profile) PrivateAPIKey() string {
	return p.GetString(privateAPIKey)
}

// SetPrivateAPIKey set configured private api key
func (p Profile) SetPrivateAPIKey(v string) {
	p.Set(privateAPIKey, v)
}

// OpsManagerURL get configured ops manager base url
func (p Profile) OpsManagerURL() string {
	return p.GetString(opsManagerURL)
}

// SetOpsManagerURL set configured ops manager base url
func (p Profile) SetOpsManagerURL(v string) {
	p.Set(opsManagerURL, v)
}

// ProjectID get configured project ID
func (p Profile) ProjectID() string {
	return p.GetString(ProjectID)
}

func (p Profile) SetProjectID(v string) {
	p.Set(ProjectID, v)
}

// Save save the configuration to disk
func Load() error {
	// Find home directory.
	configDir, err := configHome()
	if err != nil {
		return err
	}
	viper.SetConfigType(configType)
	viper.SetConfigName(Name)
	viper.SetConfigPermissions(0600)
	viper.AddConfigPath(configDir)

	viper.SetEnvPrefix(Name)
	// TODO: review why this is not working as expected
	viper.RegisterAlias(baseURL, opsManagerURL)
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		// ignore if it doesn't exists
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil
		}
		return err
	}
	return nil
}

// Save the configuration to disk
func (p Profile) Save() error {
	exists, err := afero.DirExists(p.fs, p.configDir)
	if err != nil {
		return err
	}
	if !exists {
		err := p.fs.MkdirAll(p.configDir, 0700)
		if err != nil {
			return err
		}
	}
	// TODO: We can now read but not write, see https://github.com/spf13/viper/pull/813
	configFile := fmt.Sprintf("%s/%s.toml", p.configDir, Name)
	return viper.WriteConfigAs(configFile)
}
