package main

import "github.com/nlopes/slack"

const (
	SlackToken      = "xoxb-324237408453-0yYuvLezDwcn8OclXo7jsjKq"
	RejectedChannel = "shutterstock_rejected"
	ApprovedChannel = "shutterstock_approved"
	EarningsChannel = "shutterstock"
)

type Slack interface {
	NewSlackClient() TypeSlack
	Send(msg SlackMessage, slackChannel string) error
	ErrorMessage() error
}

type TypeSlack struct {
	Logger TypeLogger
	Client *slack.Client
}

type SlackMessage struct {
	ImageURL string
	Title    string
	Text     string
}

func NewSlack() TypeSlack {
	return TypeSlack{}
}

func (s TypeSlack) NewSlackClient(logger TypeLogger) TypeSlack {
	return TypeSlack{
		Logger: logger,
		Client: slack.New(SlackToken),
	}
}
