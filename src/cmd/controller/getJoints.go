/*
Copyright © 2024 Antoine Harle antoine.harle@proton.me

*/
package controller

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getJointsCmd represents the getJoints command
var getJointsCmd = &cobra.Command{
	Use:   "getJoints",
	Short: "getJoints command descritpion",
	Long: `getJoints command descritpion`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getJoints called")
	},
}

func init() {
	controllerCmd.AddCommand(getJointsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getJointsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getJointsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
