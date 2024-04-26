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
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func generateToken(pemPath string, appID string) (string, error) {
	pemBytes, err := os.ReadFile(pemPath)
	if err != nil {
		return "", err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(pemBytes)
	if err != nil {
		return "", err
	}

	const timeout = 10 * time.Minute

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(timeout).Unix(),
		"iss": appID,
	})

	// Sign the JWT token with the private key
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

type accessTokensRequest struct {
	Repositories  []string          `json:"repositories"`
	RepositoryIDs []int             `json:"repository_ids"`
	Permissions   map[string]string `json:"permissions"`
}

type installation struct {
	ID      int `json:"id"`
	Account struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"account"`
	RepositorySelection string `json:"repository_selection"`
	AccessTokensURL     string `json:"access_tokens_url"`
	RepositoriesURL     string `json:"repositories_url"`
	HTMLURL             string `json:"html_url"`
	AppID               int    `json:"app_id"`
	AppSlug             string `json:"app_slug"`
	TargetID            int    `json:"target_id"`
	TargetType          string `json:"target_type"`
	Permissions         struct {
		Actions      string `json:"actions"`
		Contents     string `json:"contents"`
		Metadata     string `json:"metadata"`
		PullRequests string `json:"pull_requests"`
	} `json:"permissions"`
	Events                 []any     `json:"events"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
	SingleFileName         any       `json:"single_file_name"`
	HasMultipleSingleFiles bool      `json:"has_multiple_single_files"`
	SingleFilePaths        []any     `json:"single_file_paths"`
	SuspendedBy            any       `json:"suspended_by"`
	SuspendedAt            any       `json:"suspended_at"`
}

func getInstallationID(ctx context.Context, repo, token string) (int, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("https://api.github.com/repos/%s/installation", repo), nil)
	if err != nil {
		return -1, err
	}
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return -1, err
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return -1, fmt.Errorf("invalid http status code: %s", res.Status)
	}
	defer res.Body.Close()
	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return -1, err
	}
	i := installation{}
	err = json.Unmarshal(buf, &i)
	if err != nil {
		return -1, err
	}
	return i.ID, nil
}

type accessTokensResponse struct {
	Token       string    `json:"token"`
	ExpiresAt   time.Time `json:"expires_at"`
	Permissions struct {
		Issues   string `json:"issues"`
		Contents string `json:"contents"`
	} `json:"permissions"`
	RepositorySelection string `json:"repository_selection"`
	Repositories        []struct {
		ID       int    `json:"id"`
		NodeID   string `json:"node_id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		Owner    struct {
			Login             string `json:"login"`
			ID                int    `json:"id"`
			NodeID            string `json:"node_id"`
			AvatarURL         string `json:"avatar_url"`
			GravatarID        string `json:"gravatar_id"`
			URL               string `json:"url"`
			HTMLURL           string `json:"html_url"`
			FollowersURL      string `json:"followers_url"`
			FollowingURL      string `json:"following_url"`
			GistsURL          string `json:"gists_url"`
			StarredURL        string `json:"starred_url"`
			SubscriptionsURL  string `json:"subscriptions_url"`
			OrganizationsURL  string `json:"organizations_url"`
			ReposURL          string `json:"repos_url"`
			EventsURL         string `json:"events_url"`
			ReceivedEventsURL string `json:"received_events_url"`
			Type              string `json:"type"`
			SiteAdmin         bool   `json:"site_admin"`
		} `json:"owner"`
		Private          bool      `json:"private"`
		HTMLURL          string    `json:"html_url"`
		Description      string    `json:"description"`
		Fork             bool      `json:"fork"`
		URL              string    `json:"url"`
		ArchiveURL       string    `json:"archive_url"`
		AssigneesURL     string    `json:"assignees_url"`
		BlobsURL         string    `json:"blobs_url"`
		BranchesURL      string    `json:"branches_url"`
		CollaboratorsURL string    `json:"collaborators_url"`
		CommentsURL      string    `json:"comments_url"`
		CommitsURL       string    `json:"commits_url"`
		CompareURL       string    `json:"compare_url"`
		ContentsURL      string    `json:"contents_url"`
		ContributorsURL  string    `json:"contributors_url"`
		DeploymentsURL   string    `json:"deployments_url"`
		DownloadsURL     string    `json:"downloads_url"`
		EventsURL        string    `json:"events_url"`
		ForksURL         string    `json:"forks_url"`
		GitCommitsURL    string    `json:"git_commits_url"`
		GitRefsURL       string    `json:"git_refs_url"`
		GitTagsURL       string    `json:"git_tags_url"`
		GitURL           string    `json:"git_url"`
		IssueCommentURL  string    `json:"issue_comment_url"`
		IssueEventsURL   string    `json:"issue_events_url"`
		IssuesURL        string    `json:"issues_url"`
		KeysURL          string    `json:"keys_url"`
		LabelsURL        string    `json:"labels_url"`
		LanguagesURL     string    `json:"languages_url"`
		MergesURL        string    `json:"merges_url"`
		MilestonesURL    string    `json:"milestones_url"`
		NotificationsURL string    `json:"notifications_url"`
		PullsURL         string    `json:"pulls_url"`
		ReleasesURL      string    `json:"releases_url"`
		SSHURL           string    `json:"ssh_url"`
		StargazersURL    string    `json:"stargazers_url"`
		StatusesURL      string    `json:"statuses_url"`
		SubscribersURL   string    `json:"subscribers_url"`
		SubscriptionURL  string    `json:"subscription_url"`
		TagsURL          string    `json:"tags_url"`
		TeamsURL         string    `json:"teams_url"`
		TreesURL         string    `json:"trees_url"`
		CloneURL         string    `json:"clone_url"`
		MirrorURL        string    `json:"mirror_url"`
		HooksURL         string    `json:"hooks_url"`
		SvnURL           string    `json:"svn_url"`
		Homepage         string    `json:"homepage"`
		Language         any       `json:"language"`
		ForksCount       int       `json:"forks_count"`
		StargazersCount  int       `json:"stargazers_count"`
		WatchersCount    int       `json:"watchers_count"`
		Size             int       `json:"size"`
		DefaultBranch    string    `json:"default_branch"`
		OpenIssuesCount  int       `json:"open_issues_count"`
		IsTemplate       bool      `json:"is_template"`
		Topics           []string  `json:"topics"`
		HasIssues        bool      `json:"has_issues"`
		HasProjects      bool      `json:"has_projects"`
		HasWiki          bool      `json:"has_wiki"`
		HasPages         bool      `json:"has_pages"`
		HasDownloads     bool      `json:"has_downloads"`
		Archived         bool      `json:"archived"`
		Disabled         bool      `json:"disabled"`
		Visibility       string    `json:"visibility"`
		PushedAt         time.Time `json:"pushed_at"`
		CreatedAt        time.Time `json:"created_at"`
		UpdatedAt        time.Time `json:"updated_at"`
		Permissions      struct {
			Admin bool `json:"admin"`
			Push  bool `json:"push"`
			Pull  bool `json:"pull"`
		} `json:"permissions"`
		AllowRebaseMerge    bool   `json:"allow_rebase_merge"`
		TemplateRepository  any    `json:"template_repository"`
		TempCloneToken      string `json:"temp_clone_token"`
		AllowSquashMerge    bool   `json:"allow_squash_merge"`
		AllowAutoMerge      bool   `json:"allow_auto_merge"`
		DeleteBranchOnMerge bool   `json:"delete_branch_on_merge"`
		AllowMergeCommit    bool   `json:"allow_merge_commit"`
		SubscribersCount    int    `json:"subscribers_count"`
		NetworkCount        int    `json:"network_count"`
		License             struct {
			Key     string `json:"key"`
			Name    string `json:"name"`
			URL     string `json:"url"`
			SpdxID  string `json:"spdx_id"`
			NodeID  string `json:"node_id"`
			HTMLURL string `json:"html_url"`
		} `json:"license"`
		Forks      int `json:"forks"`
		OpenIssues int `json:"open_issues"`
		Watchers   int `json:"watchers"`
	} `json:"repositories"`
}

func getAccessToken(ctx context.Context, installationID int, token string, r *accessTokensRequest) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("https://api.github.com/app/installations/%d/access_tokens", installationID), nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	buf, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	req.Body = io.NopCloser(bytes.NewBuffer(buf))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return "", fmt.Errorf("invalid http status code: %s", res.Status)
	}
	defer res.Body.Close()
	buf, err = io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	i := accessTokensResponse{}
	err = json.Unmarshal(buf, &i)
	if err != nil {
		return "", err
	}
	return i.Token, nil
}

func convertPermissions(s string) map[string]string {
	result := map[string]string{}
	for _, permission := range strings.Split(s, ",") {
		data := strings.Split(permission, "=")
		result[data[0]] = data[1]
	}
	return result
}

func run(pem, appID, owner, repo string, perm map[string]string) error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	token, err := generateToken(pem, appID)
	if err != nil {
		return err
	}

	id, err := getInstallationID(ctx, fmt.Sprintf("%s/%s", owner, repo), token)
	if err != nil {
		return err
	}

	accessToken, err := getAccessToken(ctx, id, token, &accessTokensRequest{
		Repositories: []string{repo},
		Permissions:  perm,
	})
	if err != nil {
		return err
	}

	fmt.Println(accessToken)
	return nil
}

func main() {
	pem := flag.String("pem", "", "path to pem file")
	appID := flag.String("app_id", "", "id of github app")
	owner := flag.String("owner", "", "github owner")
	repo := flag.String("repo", "", "github repo")
	perm := flag.String("perm", "", "access token permissions")

	flag.Parse()

	if pem == nil || *pem == "" {
		log.Fatalf("pem flag is required")
	}

	if appID == nil || *appID == "" {
		log.Fatalf("app_id flag is required")
	}

	if owner == nil || *owner == "" {
		log.Fatalf("owner flag is required")
	}

	if repo == nil || *repo == "" {
		log.Fatalf("repo flag is required")
	}

	if perm == nil || *perm == "" {
		log.Fatalf("perm flag is required")
	}

	err := run(*pem, *appID, *owner, *repo, convertPermissions(*perm))
	if err != nil {
		log.Fatal(err)
	}
}
