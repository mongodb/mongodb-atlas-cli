// Copyright 2024 MongoDB Inc
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

package main

import (
	"context"
	"net/http"

	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/google/go-github/v61/github"
)

var (
	read  = "read"
	write = "write"
)

type apixBotRepo struct {
	owner string
	name  string
}

type apixBot struct {
	pemPath string
	appID   int64
	repo    apixBotRepo
}

func (bot apixBot) githubClient() (*github.Client, error) {
	itr, err := ghinstallation.NewAppsTransportKeyFromFile(http.DefaultTransport, bot.appID, bot.pemPath)
	if err != nil {
		return nil, err
	}

	return github.NewClient(&http.Client{Transport: itr}), nil
}

func (bot apixBot) installationID(ctx context.Context) (int64, error) {
	client, err := bot.githubClient()
	if err != nil {
		return -1, err
	}

	installation, _, err := client.Apps.FindRepositoryInstallation(ctx, bot.repo.owner, bot.repo.name)
	if err != nil {
		return -1, err
	}

	return *installation.ID, nil
}

func (bot apixBot) accessToken(ctx context.Context, installationID int64) (string, error) {
	client, err := bot.githubClient()
	if err != nil {
		return "", err
	}

	token, _, err := client.Apps.CreateInstallationToken(ctx, installationID, &github.InstallationTokenOptions{
		Permissions: &github.InstallationPermissions{
			Actions:  &read,
			Contents: &write,
		},
	})
	if err != nil {
		return "", err
	}

	return token.GetToken(), nil
}

func (bot apixBot) InstallationAccessToken(ctx context.Context) (string, error) {
	id, err := bot.installationID(ctx)
	if err != nil {
		return "", err
	}
	return bot.accessToken(ctx, id)
}
