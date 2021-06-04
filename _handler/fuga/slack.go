package main

import (
	"context"

	"github.com/ebi-yade/makefile-sandbox/utils"
	"github.com/slack-go/slack"
)

type client struct {
	slack.Client
}

var (
	slackToken = utils.MustNonEmptyEnv("SLACK_TOKEN")
)

func newClient() *client {
	return &client{
		*slack.New(slackToken),
	}
}

func (c *client) sendMessage(ctx context.Context, msg string) error {
	_, _, err := c.PostMessageContext(ctx, "#random", slack.MsgOptionText(msg, true))
	return err
}
