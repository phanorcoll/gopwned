package main

import (
	"os"
	"time"

	"bitbucket.com/phanorcoll/clipwned/lib/pwned"
	raven "github.com/getsentry/raven-go"
	"github.com/urfave/cli"
)

func init() {
	raven.SetDSN("https://778bf3fe9e124a8293467ecf0098e361:9ebed7afc7d74167af0b1972eb183592@sentry.io/115941")
}

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
	app.UsageText = "gopwned verify <email>"

	app.Commands = []cli.Command{
		{
			Name:  "verify",
			Usage: "Verifys the email account provided",
			Action: func(c *cli.Context) {
				pwnedapi.GetEmail(c.Args().First())
			},
		},
	}

	app.Run(os.Args)
}
