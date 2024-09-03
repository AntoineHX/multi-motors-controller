/*
Copyright © 2024 Antoine Harle antoine.harle@proton.me

*/
package cmd

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"
)

// setJointsCmd represents the setJoints command
var setJointsCmd = &cobra.Command{
	Use:   "setJoints",
	Short: "setJoints command descritpion",
	Long: `setJoints command descritpion`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("setJoints called with target :", cmd.Flag("vel").Value)
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
	//TODO: Anonymous flag (with anonymous flag group ?)
	//TODO: Fix input of multiple values as slice
	setJointsCmd.Flags().Float32Slice("vel", nil, "Target joint values") //or IntSlice ?
	setJointsCmd.MarkFlagRequired("vel")
}
