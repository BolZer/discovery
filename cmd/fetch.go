package cmd

import (
	"discovery/discovery"
	"discovery/discovery/mh4u"
	"discovery/discovery/mhgu"
	"discovery/discovery/mhw"
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
)

var path string

var Fetch = cli.Command{
	Name:    "fetch",
	Aliases: []string{"f"},
	Usage:   "fetch monster data for the supported games.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "path",
			Required:    true,
			Destination: &path,
			Usage:       "used to point to a target directory where the json is saved",
		},
	},
	Action: func(context *cli.Context) error {

		fmt.Println("Starting fetching monster data...")

		if _, err := os.Stat(path); os.IsNotExist(err) {
			return err
		}

		var buffer []discovery.Entity

		err := mhw.StartDiscovery(func(entities []discovery.Entity) {
			buffer = append(buffer, entities...)
		})

		if err != nil {
			return err
		}

		err = mhgu.StartDiscovery(func(entities []discovery.Entity) {
			buffer = append(buffer, entities...)
		})

		if err != nil {
			return err
		}

		err = mh4u.StartDiscovery(func(entities []discovery.Entity) {
			buffer = append(buffer, entities...)
		})

		if err != nil {
			return err
		}

		file, err := json.MarshalIndent(buffer, "", " ")

		if err != nil {
			return err
		}

		err = ioutil.WriteFile(path, file, 0644)

		if err != nil {
			return err
		}

		fmt.Println(fmt.Sprintf("Done fetching monster data. Saved file to: %s", path))

		return nil
	},
}
