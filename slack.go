package main

import (
	"github.com/nlopes/slack"
	"encoding/json"
	"fmt"
)

const (
	SlackToken      = ""
	RejectedChannel = "shutterstock_rejected"
	ApprovedChannel = "shutterstock_approved"
	EarningsChannel = "shutterstock"
)

type Slack interface {
	Send(body []byte, dataType string) error
	ErrorMessage() error
}

type TypeSlack struct {
	SlackChannel map[string]string
	Client       *slack.Client
}

type SlackMessage struct {
	ImageUrl string `json:"image_url"`
	Title    string `json:"title"`
	Text     string `json:"text"`
}

func NewSlack() TypeSlack {
	return TypeSlack{
		SlackChannel: map[string]string{
			"rejected": "shutterstock_rejected",
			"approved": "shutterstock_approved",
			"earnings": "shutterstock"},
		Client: slack.New(SlackToken),
	}
}

func (s TypeSlack) Send(body []byte, dataType string) error {

	msg, err := s.GetMessage(body)
	if err != nil {
		return err
	}

	attachment := []slack.Attachment{
		slack.Attachment{
			ImageURL: msg.ImageUrl,
			Fields: []slack.AttachmentField{
				slack.AttachmentField{
					Title: msg.Title,
				},
			},
		},
	}

	fmt.Println(attachment)

	//_, _, err = s.Client.PostMessage(
	//	s.SlackChannel[dataType],
	//	msg.Text,
	//	slack.PostMessageParameters{Attachments: attachment})
	//if err != nil {
	//	return err
	//}

	//fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)

	return nil
}

func (s TypeSlack) GetMessage(body []byte) (SlackMessage, error) {
	msg := SlackMessage{}

	if err := json.Unmarshal(body, &msg); err != nil {
		return SlackMessage{}, err
	}

	return msg, nil

}

func (s TypeSlack) ErrorMessage() error {

	return nil

}
