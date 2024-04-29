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
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
)

func run(appID int64, pem, owner, repo string) error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	bot := &apixBot{
		appID:   appID,
		pemPath: pem,
		repo: apixBotRepo{
			owner: owner,
			name:  repo,
		},
	}

	accessToken, err := bot.InstallationAccessToken(ctx)
	if err != nil {
		return err
	}

	fmt.Println(accessToken)
	return nil
}

func main() {
	pem := flag.String("pem", "", "path to pem file")
	appID := flag.Int64("app_id", -1, "id of github app")
	owner := flag.String("owner", "", "github owner")
	repo := flag.String("repo", "", "github repo")

	flag.Parse()

	if pem == nil || *pem == "" {
		log.Fatalf("pem flag is required")
	}

	if appID == nil || *appID < 0 {
		log.Fatalf("app_id flag is required")
	}

	if owner == nil || *owner == "" {
		log.Fatalf("owner flag is required")
	}

	if repo == nil || *repo == "" {
		log.Fatalf("repo flag is required")
	}

	if err := run(*appID, *pem, *owner, *repo); err != nil {
		log.Fatalln(err)
	}
}
