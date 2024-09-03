/*
Copyright © 2024 Antoine Harle antoine.harle@proton.me

*/
package controller

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve command descritpion",
	Long: `serve command descritpion`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called with: Port ", cmd.Flag("port").Value)
		fmt.Println(cmd.CommandPath())
	},
}

func init() {
	controllerCmd.AddCommand(serveCmd)
	// motorCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	serveCmd.Flags().Uint16P("port", "p", 8080, "Port")
}