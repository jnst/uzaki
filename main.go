package main

import (
	"github.com/jnst/uzaki/applestore"
	"github.com/jnst/uzaki/msg"
	"github.com/jnst/uzaki/yamatomichi"
	"log"
	"time"
)

func main() {
	//checkYamatomichi()
	checkAppleStore()
}

func checkYamatomichi() {
	for {
		url := yamatomichi.CreateURL()
		log.Printf("requesting to %s", url)

		s, err := yamatomichi.Get(url)
		if err != nil {
			log.Println(s)
		}

		yamatomichi.Print(s)

		time.Sleep(6 * time.Minute)
	}
}

func checkAppleStore() {
	for {
		url := applestore.CreateURL()
		log.Printf("requesting to %s", url)

		s, err := applestore.Get(url)
		if err != nil {
			msg.Notify("Apple Watch Series 6", err.Error(), false)
			time.Sleep(1 * time.Hour)
			continue
		}

		if applestore.Check(s) {
			msg.Notify("Apple Watch Series 6", "チャコールブレイデッドソロループ in stock now!", true)
		}

		time.Sleep(6 * time.Minute)
	}
}
