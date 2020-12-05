package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jnst/uzaki/domain"
	"github.com/jnst/uzaki/infrastructure"
	"github.com/jnst/uzaki/usecase"
)

func main() {
	lambda.Start(func() {
		var n domain.Notifier = &infrastructure.Slack{}
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
