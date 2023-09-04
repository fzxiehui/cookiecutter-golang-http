package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/config"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/log"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/http"
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
		cfg := config.ConfigViper()
		log.Debug("loglevel:", cfg.GetString("loglevel"))
		servers, cleanup, err := newHttpServe(cfg)
		if err != nil {
			panic(err)
		}
		http.Run(servers.ServerHTTP, fmt.Sprintf(":%d", cfg.GetInt("http.port")))
		defer cleanup()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVarP(&configPath, "config", "c", "", "config file path")
}
