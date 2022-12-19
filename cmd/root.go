package cmd

import (
	"github.com/bzp2010/schedule/internal/config"
	"github.com/bzp2010/schedule/internal/database"
	"github.com/bzp2010/schedule/internal/log"
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

	cmd.PersistentFlags().StringVarP(&configFile, "config", "c", "config/config.yaml", "config file")

	return cmd
}

func run(cmd *cobra.Command, args []string) error {
	// config
	cfg := config.NewDefaultConfig()
	if err := config.SetupConfig(&cfg, configFile); err != nil {
		return errors.Wrap(err, "failed to setup config")
	}

	// logger
	if err := log.SetupLogger(cfg); err != nil {
		return errors.Wrap(err, "failed to setup logger")
	}

	//database
	if err := database.SetupDatabase(cfg); err != nil {
		return errors.Wrap(err, "failed to setup database")
	}

	return nil
}
