package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "boom"
	app.Usage = "make an explosive entrance"
	app.Version = "0.0.1"

	os.Setenv("SAMPLE_ENV", "sample env dayo")

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang, l",
			Value: "english",
			Usage: "language for the greeting",
		},
		cli.StringFlag{
			Name:  "meridian, m",
			Value: "AM",
			Usage: "meridian ofr the greeting",
		},
		cli.StringFlag{
			Name:  "time, t",
			Value: "07:00",
			Usage: "`your time` for the greeting",
		},
		cli.StringFlag{
			Name:   "aaa, a",
			Value:  "sample",
			EnvVar: "SAMPLE_ENV",
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Println("--Action--")

		fmt.Printf("c.NArg()		: %+v\n", c.NArg())
		fmt.Printf("c.Args()		: %+v\n", c.Args())
		fmt.Printf("c.Args().get(0)	: %+v\n", c.Args().Get(0))
		fmt.Printf("c.FlagNames		: %+v\n", c.FlagNames)
		return nil
	}

	app.Run(os.Args)
}
