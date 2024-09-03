/*
Copyright © 2024 Antoine Harle antoine.harle@proton.me

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// moveVelCmd represents the moveVel command
var moveVelCmd = &cobra.Command{
	Use:   "moveVel",
	Short: "moveVel command descritpion",
	Long: `moveVel command descritpion`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("moveVel called")
	},
}

func init() {
	motorCmd.AddCommand(moveVelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// moveVelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// moveVelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
