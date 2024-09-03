/*
Copyright © 2024 Antoine Harle antoine.harle@proton.me

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// setJointsCmd represents the setJoints command
var setJointsCmd = &cobra.Command{
	Use:   "setJoints",
	Short: "setJoints command descritpion",
	Long: `setJoints command descritpion`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("setJoints called")
	},
}

func init() {
	controllerCmd.AddCommand(setJointsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setJointsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setJointsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
