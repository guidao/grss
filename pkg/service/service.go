package service

import (
	"context"
	"strings"

	"github.com/google/go-github/v39/github"
	"github.com/gorilla/feeds"
	"github.com/guidao/grss/config"
	"golang.org/x/oauth2"
)

type GRSSService struct {
	github *github.Client
}

func NewService() *GRSSService {
	token := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: config.GetConf().Github.Token,
	})

	client := oauth2.NewClient(context.Background(), token)
	gh := github.NewClient(client)

	return &GRSSService{
		github: gh,
	}
}

func (r *GRSSService) FetchGithub() (string, error) {
	feed := &feeds.Feed{
		Title:       "github",
		Link:        &feeds.Link{Href: "http://github.com/guidao"},
		Description: "github",
	}
	conf := config.GetConf()
	if conf.Github == nil {
		return "", nil
	}
	for _, repo := range conf.Github.Repos {
		fields := strings.Split(repo, "/")
		issues, _, err := r.github.Activity.ListIssueEventsForRepository(context.Background(), fields[0], fields[1], &github.ListOptions{})
		if err != nil {
			return "", err
		}

		for _, issue := range issues {
			feed.Items = append(feed.Items, &feeds.Item{
				Title:       "[" + repo + "] " + issue.Issue.GetTitle(),
				Link:        &feeds.Link{Href: issue.Issue.GetHTMLURL()},
				Description: repo,
				Created:     issue.Issue.GetUpdatedAt(),
				Content:     issue.Issue.GetBody(),
			})
		}
	}

	return feed.ToAtom()
}
