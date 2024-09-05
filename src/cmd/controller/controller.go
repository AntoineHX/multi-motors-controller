/*
Copyright © 2024 Antoine Harle antoine.harle@proton.me

*/
package controller

import (
	"fmt"
	// "log"

	"github.com/AntoineHX/multi-motors-controller/src/cmd"
	motor "github.com/AntoineHX/multi-motors-controller/src/cmd/motor"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var(
	motor_configs []motor.Config //Configs of the motors
)

//Cobra CLI
// controllerCmd represents the controller command
var controllerCmd = &cobra.Command{
	Use:   "controller",
	Short: "Motor controller command",
	Long: `Motor controller command`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help() //Nothing to do, show help
	},
}

func init() {
	cmd.RootCmd.AddCommand(controllerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// controllerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// controllerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Update motor configs from config file
func updateConfig(){
	var curr_config motor.Config
	i := 0
	config_id := fmt.Sprintf("motors.%d.id",i)
	for viper.IsSet(config_id){
		// log.Printf("%v / %v", viper.GetInt(config_id), int(motorID))
		curr_config = motor.Config{}
		motor.ExtractConfig(i, &curr_config)
		motor_configs = append(motor_configs,curr_config)
		//Next motor
		i++
		config_id = fmt.Sprintf("motors.%d.id",i)
	}

	// log.Printf("Using config: %+v", motor_configs)
}