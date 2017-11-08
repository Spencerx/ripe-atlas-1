package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"strconv"
)

// init injects our "ip" related commands/options.
func init() {
	// Fill-in the various commands
	cliCommands = append(cliCommands, cli.Command{
		Name:        "ip",
		Usage:       "returns current ip",
		Description: "shorthand for getting current ip",
		Action:      cmdIP,
	})
}

// shortcuts

// cmdIP is a short for displaying the IPs for one probe
func cmdIP(c *cli.Context) error {

	var (
		probeID int
	)

	args := c.Args()
	if len(args) == 1 {
		probeID, _ = strconv.Atoi(args[0])
	}

	if probeID == 0 {
		if mycnf.DefaultProbe == 0 {
			log.Fatal("Error: you must specify a probe ID!")
		} else {
			probeID = mycnf.DefaultProbe
		}
	}

	p, err := client.GetProbe(probeID)
	if err != nil {
		fmt.Printf("err: %v", err)
		os.Exit(1)
	}

	fmt.Printf("IPv4: %s IPv6: %s\n", p.AddressV4, p.AddressV6)
	return nil
}
