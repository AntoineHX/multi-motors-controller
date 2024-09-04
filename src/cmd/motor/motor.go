/*
Copyright © 2024 Antoine Harle antoine.harle@proton.me

*/
package motor

import (
	"fmt"
	"log"

	"github.com/AntoineHX/multi-motors-controller/src/cmd"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

//Cobra CLI
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

// Update curr_config from config file
func updateConfig()(bool){
	//Find correct ID in config file
	i := 0
	config_id := fmt.Sprintf("motors.%d.id",i)
	for viper.IsSet(config_id){
		// log.Printf("%v / %v", viper.GetInt(config_id), int(motorID))
		if viper.GetInt(config_id) == int(motorID) {
			//Extract config for this motor
			if !ExtractConfig(i, &curr_config) { 
				//Extraction failed->stop
				return false
			}
			break
		}
		//Next motor
		i++
		config_id = fmt.Sprintf("motors.%d.id",i)
	}
	//Sanity check
	if curr_config.Id != motorID {
		log.Fatalf("Failed to update config. Requested ID: %d. Got: %d.", motorID, curr_config.Id)
		return false
	}
	log.Printf("Using config: %+v", curr_config)
	return true
}

func ExtractConfig(idx int, config *Config)(bool){
	//Extract config for this motor
	err := viper.UnmarshalKey(fmt.Sprintf("motors.%d",idx), config) //&
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
		return false
	}
	return true
}