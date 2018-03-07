package main

import (
	"log"

	"github.com/jeremyliu/staccato/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
