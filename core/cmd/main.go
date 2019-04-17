package main

import (
	"fmt"
	"github.com/unarxiv/adbt/pkg"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Name = "adbt"
	app.Usage = "Automatic Database Backup Toolkit"
	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"add"},
			Usage:   "[Wizard] Add a backup job",
			Action:  addHandler,
		},
		{
			Name:    "install",
			Aliases: []string{"i"},
			Usage:   "Install daemon to the System",
			Action:  func(c *cli.Context) error {
				fmt.Println("Installing...")
				fmt.Println("It may ask for root/Administrator privilege.")
				adbt.InstallService()
				fmt.Println("Installed!")
				return nil
			},
		},
		{
			Name:    "uninstall",
			Aliases: []string{"ui"},
			Usage:   "Remove daemon to the System",
			Action:  func(c *cli.Context) error {
				fmt.Println("Uninstalling...")
				fmt.Println("It may ask for root/Administrator privilege.")
				adbt.UninstallService()
				fmt.Println("Uninstalled!")
				return nil
			},
		},
		{
			Name:    "serve",
			Aliases: []string{"s"},
			Usage:   "Start the web server [FOR DEBUG ONLY]",
			Action:  func(c *cli.Context) error {
				adbt.NewServer(adbt.DaemonPort)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}