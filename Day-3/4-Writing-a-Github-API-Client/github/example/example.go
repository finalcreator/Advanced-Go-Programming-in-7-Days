package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Advanced-Go/Day-3/4-Writing-a-Github-API-Client/github"
)

var apiToken = os.Getenv("GITHUB_API_TOKEN")

func main() {
	ctx := context.Background()
	c := github.NewGithubClient(ctx, apiToken)
	repos, _, err := c.Repositories.List(ctx, "theodesp")
	if err != nil {
		fmt.Println(err)
	}

	for _, repo := range repos {
		fmt.Println(repo)
	}
}
