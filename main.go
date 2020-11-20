package main

import (
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jnst/uzaki/applestore"
	"github.com/jnst/uzaki/msg"
	"github.com/jnst/uzaki/yamatomichi"
)

func main() {
	lambda.Start(checkAppleStore)
}

func checkYamatomichi() {
	for {
		url := yamatomichi.CreateURL()
		log.Printf("requesting to %s\n", url)

		s, err := yamatomichi.Get(url)
		if err != nil {
			log.Println(s)
		}

		yamatomichi.Print(s)

		time.Sleep(6 * time.Minute)
	}
}

func checkAppleStore() error {
	url := applestore.CreateURL()
	fmt.Printf("requesting to %s\n", url)

	s, err := applestore.Get(url)
	if err != nil {
		return err
	}

	if applestore.Check(s) {
		fmt.Println("in stock now.")
		msg.Notify("Apple Watch Series 6", "チャコールブレイデッドソロループ in stock now!", true)
	} else {
		fmt.Println("nothing update.")
	}

	return nil
}
