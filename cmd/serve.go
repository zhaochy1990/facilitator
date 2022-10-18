/*
Package cmd
Copyright © 2022 Chaoyi Zhao
*/
package cmd

import (
	"fmt"
	"github.com/zhaochy1990/facilitator/server"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the server and serves the HTTP REST and gRPC APIs",
	Long: `This command opens the network ports and listens to HTTP and gRPC API requests.

## Configuration

facilitator can be configured using environment variables as well as a configuration file.
For more information on configuration options, open the configuration documentation:

>> https://github.com/zhaochy1990/facilitator/config.yml <<`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("serve called, config path: [%s]\n", configpath)

		if configpath == "" {
			return server.ServeAll(cmd.Context(), nil)
		}

		abs := filepath.IsAbs(configpath)
		if !abs {
			dir, _ := os.Getwd()
			configpath = path.Join(dir, configpath)
			fmt.Println(configpath)
		}

		return server.ServeAll(cmd.Context(), &configpath)
	},
}

var configpath string

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	serveCmd.PersistentFlags().StringVarP(&configpath, "config", "c", "", "config file path")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
