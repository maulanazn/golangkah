package main

import (
	"fmt"
	"golangkah/greetings"
	"log"
)

func main() {
	log.SetPrefix("greetings log: ")
	log.SetFlags(0)

	names := []string{"maulana", "dimas", "fatih"}

	// message, err := greetings.Hello(names)
	message, err := greetings.HelloToPeople(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
