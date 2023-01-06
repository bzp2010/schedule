package cmd

import (
	"fmt"

	"github.com/bzp2010/schedule/internal/database"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	migrateConfirm bool
)

func newMigrateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Perform a database migration to generate or update tables",
		RunE:  runMigrate,
	}

	cmd.Flags().BoolVarP(&migrateConfirm, "yes", "y", false, "confirm run migration")

	return cmd
}

func runMigrate(cmd *cobra.Command, args []string) error {
	if !migrateConfirm {
		fmt.Println(`
	If you need to perform a migration, you need to make sure that you have backed up your data.

	After that, you need to add the --yes flag for confirmation.
		`)
		return nil
	}

	_, err := setup()
	if err != nil {
		return err
	}

	// perform migrate
	err = database.Migrate(database.GetDatabase())
	if err != nil {
		return errors.Wrap(err, "failed to migrate database model")
	}

	return nil
}
