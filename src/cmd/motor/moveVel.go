/*
Copyright © 2024 Antoine Harle antoine.harle@proton.me

*/
package motor

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
		fmt.Println("moveVel called with: ID ", cmd.Flag("id").Value, " Vel ", cmd.Flag("vel").Value)
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
	//TODO: Anonymous flag (with anonymous flag group ?)
	moveVelCmd.Flags().Float32("vel", 0, "Velocity (degrees/s)")
}
