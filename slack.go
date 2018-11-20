package main

import "github.com/nlopes/slack"

const (
	SlackToken      = ""
	RejectedChannel = "shutterstock_rejected"
	ApprovedChannel = "shutterstock_approved"
	EarningsChannel = "shutterstock"
)

type Slack interface {
	Send(msg SlackMessage, slackChannel string) error
	ErrorMessage() error
}

type TypeSlack struct {
	Client *slack.Client
}

type SlackMessage struct {
	ImageURL string
	Title    string
	Text     string
}

func NewSlack() TypeSlack {
	return TypeSlack{
		Client: slack.New(SlackToken),
	}
}
