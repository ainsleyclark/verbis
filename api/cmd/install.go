package cmd

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/database/seeds"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/spf13/cobra"
)

var (
	installCmd = &cobra.Command{
		Use:   "install",
		Short: "Install will run the doctor command and then run database schema and insert any data dependant on Verbis.",
		Long:  `This command will install first run Verbis doctor to see if the database,
exists and is passable. Install will then run the migration to insert into the schema.
Seeds are also run, inserting options and any necessary configuration into the 
database.`,
		Run: func(cmd *cobra.Command, args []string) {

			// Run doctor
			db, err := doctor()
			if err != nil {
				printError(err.Error())
			}

			// Start the spinner
			printSpinner("Installing Verbis...")

			// Install the database
			if err := db.Install(); err != nil {
				fmt.Println(err)
				printError(fmt.Sprintf("A database with the name %s has already been installed. \nPlease run verbis uninstall if you want to delete it.", environment.GetDatabaseName()))
			}

			// Init Config
			con, err := config.New()
			if err != nil {
				printError(errors.Message(err))
			}

			// Set up stores & pass the database.
			store := models.New(db, *con)
			if err != nil {
				printError(err.Error())
			}

			// Run the seeds
			seeder := seeds.New(db.Sqlx, store)
			if err := seeder.Seed(); err != nil {
				printError(err.Error())
			}

			// Print success
			printSuccess("Successfully installed verbis")

			return
		},
	}
)