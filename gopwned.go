package main

import (
	"os"
	"time"

	"bitbucket.com/phanorcoll/clipwned/lib/pwned"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Version = "1.0.0"
	app.Name = "gopwned"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name: "Phanor Coll",
		},
	}
	app.Copyright = "(c) 2016 Phanor Coll"
	app.Usage = "Check to see if an email account has been compromised in a data breach!"
	app.UsageText = "gopwned verify <email> \n   gopwned company <name>"

	app.Commands = []cli.Command{
		{
			Name:      "verify",
			Usage:     "Verifys the email account provided",
			ArgsUsage: "<email-address>",
			Action: func(c *cli.Context) {
				pwned.GetBreaches(c.Args().First())
			},
		},
		{
			Name:      "company",
			Usage:     "Gets complete information about the breach",
			ArgsUsage: "<name>",
			Action: func(c *cli.Context) {
				pwned.GetBreachData(c.Args().First())
			},
		},
	}

	app.Run(os.Args)
}
