package main

import (
	"github.com/jnst/uzaki/applestore"
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
			log.Println(s)
		}

		applestore.Print(s)

		time.Sleep(6 * time.Minute)
	}
}
