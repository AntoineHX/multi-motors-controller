/*
Copyright © 2024 Antoine Harle antoine.harle@proton.me

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// motorCmd represents the motor command
var motorCmd = &cobra.Command{
	Use:   "motor",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("motor called")
	},
}

func init() {
	rootCmd.AddCommand(motorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// motorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// motorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
