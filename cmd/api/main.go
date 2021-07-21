package main

import (
	"api-testing/cmd/api/boot"
	"log"
)

func main() {
	if err := boot.Run(); err != nil {
		log.Fatal(err)
	}
}