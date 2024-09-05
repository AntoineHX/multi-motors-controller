/*
Copyright © 2024 Antoine Harle antoine.harle@proton.me

*/
package motor

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/AntoineHX/multi-motors-controller/src/proto"
)

//Cobra CLI
// moveVelCmd represents the moveVel command
var moveVelCmd = &cobra.Command{
	Use:   "moveVel",
	Short: "Set velocity of motor",
	Long: `Request motor server to move at given velocity`,
	Run: func(cmd *cobra.Command, args []string) {
		//Parse target velocity
		if len(args)!=1 {
			log.Fatalf("Wrong number of arguments")
		}else{
			vel, err :=  strconv.ParseFloat(args[0],64)
			if err!= nil{
				log.Fatalf("Failed to parse joint velocity %v", err)
			} else {
				updateConfig()
				moveVel(vel) //Set motor velocity
			}
		}
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
	moveVelCmd.Flags().Float64("", 0, "Velocity (degrees/s)") //Only used for help message
	// moveVelCmd.MarkFlagRequired("vel")
}

//gRPC Client
func moveVel(cmd_vel float64){
	// Set up a connection to the server.
	var addr = fmt.Sprintf("%s:%d", ip, curr_config.Port) //Defined in motor
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMotorClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.SetVelocity(ctx, &pb.Velocity{Velocity: cmd_vel})
	if err != nil {
		log.Fatalf("could not send: %v", err)
	} else {
		log.Printf("Requested velocity [%f]°/s to motor %d", cmd_vel, curr_config.Id)
	}
}