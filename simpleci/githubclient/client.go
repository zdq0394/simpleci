package githubclient

import (
	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var clients = make(map[string]*github.Client, 0)

func createClient(accessToken string) (c *github.Client) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	c = github.NewClient(tc)
	clients[accessToken] = c
	return
}

func GetClient(accessToken string) (c *github.Client) {
	if clients[accessToken] == nil {
		createClient(accessToken)
	}
	c = clients[accessToken]
	return
}
