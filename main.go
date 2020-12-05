package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jnst/uzaki/domain"
	"github.com/jnst/uzaki/infrastructure"
	"github.com/jnst/uzaki/usecase"
)

func main() {
	lambda.Start(func() {
		webhookURL := os.Getenv("SLACK_WEBHOOK_URL")
		fmt.Printf("Slack WebhookURL: %s\n", webhookURL)

		var n domain.Notifier = infrastructure.NewSlack(webhookURL)
		var c domain.StockChecker = &usecase.AppleWatchUsecase{}

		ok, err := c.CheckStock()
		if err != nil {
			fmt.Println(err.Error())
			n.Notify(false, "error", err.Error())
		}
		if ok {
			fmt.Println("in stock now.")
			n.Notify(true, "Apple Watch Series 6", c.String())
		} else {
			fmt.Println("nothing update.")
		}
	})
}
