/*
Copyright © 2024 Antoine Harle antoine.harle@proton.me

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// controllerCmd represents the controller command
var controllerCmd = &cobra.Command{
	Use:   "controller",
	Short: "controller command description",
	Long: `controller command description`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("controller called")
	},
}

func init() {
	rootCmd.AddCommand(controllerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// controllerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// controllerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
