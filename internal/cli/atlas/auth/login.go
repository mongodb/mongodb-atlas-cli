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

package auth

import (
	"fmt"
	"net/http"

	"github.com/mongodb/mongocli/internal/cli/require"
	"github.com/spf13/cobra"
)

func Builder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Authenticate me.",
		Annotations: map[string]string{
			"toc": "true",
		},
	}
	cmd.AddCommand(
		LoginBuilder(),
	)

	return cmd
}

type loginOpts struct {
}

func (o *loginOpts) Run() error {
	const clientID = "0oadn4hoajpzxeSEy357"
	device, err := RequestCode(
		http.DefaultClient,
		"http://localhost:8080/device/generate",
		clientID,
		[]string{"openid"},
	)
	if err != nil {
		return err
	}
	fmt.Printf("Copy this: %v\n", device.UserCode)
	fmt.Printf("Visit the URL for the auth dialog: %v\n", device.VerificationURI)
	accessToken, err := PollToken(http.DefaultClient, "http://localhost:8080/device/token", clientID, device)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Access token: %s\n", accessToken.AccessToken)
	return nil
}

func LoginBuilder() *cobra.Command {
	opts := &loginOpts{}
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Let me in.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
		Annotations: map[string]string{
			"toc": "true",
		},
		Args: require.NoArgs,
	}

	return cmd
}
