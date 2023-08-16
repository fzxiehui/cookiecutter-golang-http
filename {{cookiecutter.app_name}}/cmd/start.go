package cmd

import (
	"github.com/spf13/cobra"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/config"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/log"
)

var configPath string

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the server",
	Long:  `Starts the server`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("Starting server...")
		if configPath != "" {
			err := config.ReadViperConfigFromFile(configPath)
			if err != nil {
				log.Fatal(err)
				return
			}
		}
		cfg := config.Config()
		log.Debug("loglevel:", cfg.GetString("loglevel"))
		err := Setup()
		if err != nil {
			log.Fatal(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVarP(&configPath, "config", "c", "", "config file path")
}
