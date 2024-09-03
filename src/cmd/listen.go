/*
Copyright © 2024 Antoine Harle antoine.harle@proton.me

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listenCmd represents the listen command
var listenCmd = &cobra.Command{
	Use:   "listen",
	Short: "listen command descritpion",
	Long: `listen command descritpion`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listen called")
	},
}

func init() {
	motorCmd.AddCommand(listenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
