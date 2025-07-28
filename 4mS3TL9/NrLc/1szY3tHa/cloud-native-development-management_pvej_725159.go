// cmd/nextpulse/main.go
package main

import (
	"flag"
	"log"
	"os"

	"nextpulse/internal/nextpulse"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := nextpulse.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
