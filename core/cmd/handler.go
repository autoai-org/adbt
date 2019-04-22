package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/unarxiv/adbt/pkg"
	"github.com/urfave/cli"
	"log"
)

func addHandler(c *cli.Context) error {
	prompt := promptui.Prompt{
		Label:    "Name",
	}
	name, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}
	selectPrompt := promptui.Select{
		Label: "Database",
		Items: adbt.SupportedDatabases,
	}
	_, selectedDB, err := selectPrompt.Run()
	if err != nil {
		log.Fatal(err)
	}
	prompt = promptui.Prompt{
		Label:    "Connection String",
	}
	connectionString, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}
	selectPrompt = promptui.Select{
		Label: "Period",
		Items: adbt.SupportedPeriods,
	}
	_, selectedPeriods, err := selectPrompt.Run()
	if err != nil {
		log.Fatal(err)
	}
	prompt = promptui.Prompt{
		Label:    "Time [HH:mm], e.g 16:10",
	}
	time, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}
	isSuccessful := createJob(selectedDB,connectionString, name, selectedPeriods, time)
	if isSuccessful {
		fmt.Println("Successfully added the job")
	} else {
		fmt.Println("Failed to add the job")
	}
	return nil
}

func listHandler () {

}

func loginHandler (c *cli.Context) error {
	username := c.Args().Get(0)
	uid := c.Args().Get(1)
	adbt.Login(username, uid)
	return nil
}