package cmd

import (
	"sync"

	"github.com/bzp2010/schedule/internal/config"
	"github.com/bzp2010/schedule/internal/database"
	"github.com/bzp2010/schedule/internal/log"
	"github.com/bzp2010/schedule/internal/scheduler"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	configFile = "config/config.yaml"
)

// NewRootCommand for main package
func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "schedule",
		Short: "Timed task scheduler",
		RunE:  run,
	}

	cmd.AddCommand(newMigrateCommand())

	cmd.PersistentFlags().StringVarP(&configFile, "config", "c", "config/config.yaml", "config file")

	return cmd
}

func run(cmd *cobra.Command, args []string) error {
	err := setup()
	if err != nil {
		return err
	}

	// scheduler
	err = scheduler.SetupScheduler()
	if err != nil {
		return errors.Wrap(err, "failed to setup scheduler")
	}

	// blocking here
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()

	return nil
}

func setup() error {
	// config
	cfg := config.NewDefaultConfig()
	if err := config.SetupConfig(&cfg, configFile); err != nil {
		return errors.Wrap(err, "failed to setup config")
	}

	// logger
	if err := log.SetupLogger(cfg); err != nil {
		return errors.Wrap(err, "failed to setup logger")
	}

	// database
	if err := database.SetupDatabase(cfg); err != nil {
		return errors.Wrap(err, "failed to setup database")
	}

	return nil
}
