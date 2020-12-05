package infrastructure

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/slack-go/slack"
)

type Slack struct {
	webhookURL string
}

func NewSlack(webhookURL string) *Slack {
	return &Slack{webhookURL: webhookURL}
}

func (s *Slack) Notify(success bool, title, text string) {
	color := "good"
	if !success {
		color = "warning"
	}

	attachment := slack.Attachment{
		Color:         color,
		AuthorName:    title,
		AuthorSubname: "by uzaki",
		Text:          text,
		Ts:            json.Number(strconv.FormatInt(time.Now().Unix(), 10)),
	}
	msg := slack.WebhookMessage{
		Attachments: []slack.Attachment{attachment},
	}

	err := slack.PostWebhook(s.webhookURL, &msg)
	if err != nil {
		fmt.Println(err)
	}
}
