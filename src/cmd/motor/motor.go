/*
Copyright © 2024 Antoine Harle antoine.harle@proton.me

*/
package motor

import (
	"fmt"

	"github.com/AntoineHX/multi-motors-controller/src/cmd"

	"github.com/spf13/cobra"
)

type Config struct {
	Id uint16
	Port uint16 
	Min_pos float64
	Max_pos float64
	Max_vel float64
	Accel float64
}

type State struct {
	Angle float64 
	Velocity float64
	Error string //Error message if any
}

var(
	motorID uint16 = 0 //Requested ID
	ip string = "localhost" //localhost=127.0.0.1
	curr_config Config //Current config of the motor
)

// motorCmd represents the motor command
var motorCmd = &cobra.Command{
	Use:   "motor",
	Short: "motor command descritpion",
	Long: `motor command descritpion`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("motor called with ID: ", cmd.Flag("id").Value)
	},
}

func init() {
	cmd.RootCmd.AddCommand(motorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	motorCmd.PersistentFlags().Uint16Var(&motorID, "id", 0, "Identifier number")
	motorCmd.MarkFlagRequired("id")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// motorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
