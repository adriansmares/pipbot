package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
)

const (
	organization = "TheThingsNetwork"
	repository   = "lorawan-stack"

	perPage = 100
)

func newClient(ctx context.Context) (*github.Client, error) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GH_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc), nil
}

func run() error {
	ctx := context.Background()

	client, err := newClient(ctx)
	if err != nil {
		return err
	}

	issues := []*github.Issue{}

	if _, err := os.Stat("issues.json"); os.IsNotExist(err) {
		page := 0
		for {
			iss, _, err := client.Issues.ListByRepo(ctx, organization, repository, &github.IssueListByRepoOptions{
				State: "all",
				Sort:  "created",
				ListOptions: github.ListOptions{
					Page:    page,
					PerPage: perPage,
				},
			})
			if err != nil {
				return err
			}
			if len(iss) < perPage {
				break
			}

			issues = append(issues, iss...)
			page++

			fmt.Printf("Loaded %v issues, total %v\n", len(iss), len(issues))
		}

		b, err := json.Marshal(issues)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile("issues.json", b, 0644)
		if err != nil {
			return err
		}
	} else {
		b, err := ioutil.ReadFile("issues.json")
		if err != nil {
			return err
		}

		err = json.Unmarshal(b, &issues)
		if err != nil {
			return err
		}
	}

	g, err := createGraph(issues)
	if err != nil {
		return err
	}
	_ = g

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
