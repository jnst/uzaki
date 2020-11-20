package msg

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/slack-go/slack"
)

// Notify notifies message with slack.
func Notify(name, text string, ok bool) {
	color := "good"
	if !ok {
		color = "warning"
	}

	attachment := slack.Attachment{
		Color:         color,
		AuthorName:    name,
		AuthorSubname: "by uzaki",
		Text:          text,
		Ts:            json.Number(strconv.FormatInt(time.Now().Unix(), 10)),
	}
	msg := slack.WebhookMessage{
		Attachments: []slack.Attachment{attachment},
	}

	err := slack.PostWebhook("https://hooks.slack.com/services/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", &msg)
	if err != nil {
		fmt.Println(err)
	}
}
