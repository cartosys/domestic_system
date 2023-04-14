package main

import (
	"fmt"
	"log"

	"github.com/charmbracelet/wish"
)

func main() {
	server, err := wish.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Listening on %s...\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
