package main

import (
	"flag"
	"log"
)

func main() {
	name := flag.String("name", "", "Your Name")
	flag.Parse()

	if len(*name) != 0 {
		log.Printf("Hello %s", *name)
	} else {
		log.Fatal("Please specify your name using -name")
	}
}
