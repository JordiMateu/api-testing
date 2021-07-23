package main

import (
	boot2 "api-testing/api/boot"
	"log"
)

func main() {
	if err := boot2.Run(); err != nil {
		log.Fatal(err)
	}
}
