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

	app.Action = func(c *cli.Context) error {
		fmt.Println("--Action--")

		fmt.Print("c.NArg()		: %+v\n", c.NArg())
		return nil
	}

	app.Run(os.Args)
}
