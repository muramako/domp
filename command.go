package main

import (
	"fmt"
	"os"
	"github.com/codegangsta/cli"
)

var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		Name: "query,q",
		Value: "text",
		Usage: "query. can be used html/text/outerhtml/attrs/attr@xxx, seperator \"|\". e.g. --query \"html|attr@id\".",
	},
	cli.StringFlag{
		Name: "output,o",
		Value: "text",
		Usage: "output. can be used text/csv/json, e.g. --output \"csv\".",
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
