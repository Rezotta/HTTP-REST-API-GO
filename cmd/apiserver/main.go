package main

import (
	"log"
)

func main() {
	s := APIServer.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
