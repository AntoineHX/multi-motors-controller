/*
Copyright © 2024 Antoine Harle antoine.harle@proton.me

*/
package motor

import (
	"fmt"

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
	Short: "moveVel command descritpion",
	Long: `moveVel command descritpion`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("moveVel called with: ID ", cmd.Flag("id").Value, " Vel ", cmd.Flag("vel").Value)
		updateConfig()
		var vel, err = cmd.Flags().GetFloat64("vel")
		if err!= nil {
			log.Fatalf("Failed read requested velocity %v", err)
		}else{
			moveVel(vel)
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
	//TODO: Anonymous flag (with anonymous flag group ?)
	moveVelCmd.Flags().Float64("vel", 0, "Velocity (degrees/s)")
}

//gRPC Client
func moveVel(cmd_vel float64){
	// Set up a connection to the server.
	var addr = fmt.Sprintf("%s:%d", ip, curr_config.Port) //Defined in controller/serve
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
	}
}