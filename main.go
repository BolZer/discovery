package main

import (
	"discovery/cmd"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "discovery",
		Usage: "Scrap monster data for the Games Monster Hunter Worlds, GU and 4U",
		Commands: []*cli.Command{
			&cmd.Fetch,
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
