package main

import (
	"log"
	"time"

	"github.com/jnst/uzaki/yamatomichi"
)

func main() {
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
