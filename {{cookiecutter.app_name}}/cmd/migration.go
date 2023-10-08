package cmd

import (
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/config"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/log"
	"github.com/spf13/cobra"
)

var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "migration",
	Long:  `migration`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("Starting server...")
		if configPath != "" {
			err := config.ReadViperConfigFromFile(configPath)
			if err != nil {
				log.Fatal(err)
				return
			}
		}
		cfg := config.ConfigViper()
		log.Debug("loglevel:", cfg.GetString("loglevel"))

		mig, cleanup, err := newMigrate(cfg, args)
		if err != nil {
			panic(err)
		}
		mig.Run()
		defer cleanup()
	},
}

func init() {
	rootCmd.AddCommand(migrationCmd)
	migrationCmd.Flags().StringVarP(&configPath, "config", "c", "", "config file path")
	migrationCmd.Flags().StringSlice("strings", []string{}, "init SQL File")
}
