package main

import (
	"dbm/greetings"
	"fmt"
	"log"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(1)

	message, err := greetings.Hello("Fulano")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
