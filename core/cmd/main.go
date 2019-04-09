package main

import (
	"github.com/tucnak/climax"
)

func main() {
	demo := climax.New("dbtools")
	demo.Brief = "DBTools is an automatic backup/restore toolkit for various databases"
	demo.Version = "alpha"

	joinCmd := climax.Command{
		Name:  "backup",
		Brief: "Backup a specific database",
		Usage: `[-h=] "Connection String to the database
				[-d=] "Database Name"
				[-c=] "Column Name"`,
		Help: `More?`,

		Examples: []climax.Example{
			{
				Usecase:     `-h . "google" "com"`,
				Description: `Results in "google.com"`,
			},
		},

		Handle: backupHandler(),
	}

	demo.AddCommand(joinCmd)
	demo.Run()

}
