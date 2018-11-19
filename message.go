package main

func (s TypeSlack) Send(msg SlackMessage, slackChannel string) error {

	//attachment := []slack.Attachment{
	//	slack.Attachment{
	//		ImageURL: msg.ImageURL,
	//		Fields: []slack.AttachmentField{
	//			slack.AttachmentField{
	//				Title: msg.Title,
	//			},
	//		},
	//	},
	//}
	//
	//channelID, timestamp, err := s.Client.PostMessage(
	//	slackChannel,
	//	msg.Text,
	//	slack.PostMessageParameters{Attachments: attachment})
	//if err != nil {
	//	return err
	//}
	//
	//fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)

	return nil
}


func (s TypeSlack) ErrorMessage() error {

	return nil

}