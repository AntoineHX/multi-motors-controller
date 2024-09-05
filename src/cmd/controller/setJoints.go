/*
Copyright © 2024 Antoine Harle antoine.harle@proton.me

*/
package controller

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
// setJointsCmd represents the setJoints command
var setJointsCmd = &cobra.Command{
	Use:   "setJoints",
	Short: "Set target joint angles for the motors",
	Long: `Request the Motor Controller to command the motors to reach the target joint angles`,
	Run: func(cmd *cobra.Command, args []string) {
		//Parse target angles
		var tgt_angles []float64
		for _, arg := range args {
			angle, err := strconv.ParseFloat(arg, 64)
			if err!= nil{
				log.Fatalf("Failed to parse joint angle %v", err)
			}else{
				tgt_angles = append(tgt_angles, angle)
			}
		}
		//Set joint angles
		SetJoints(tgt_angles)
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
	setJointsCmd.Flags().Float64Slice("", nil, "Target joint values") //Only used for help message
	// setJointsCmd.MarkFlagRequired("")
}

//gRPC Client
func SetJoints(tgt_angles []float64){
	// Set up a connection to the server.
	var addr = fmt.Sprintf("%s:%d", ip, port) //Defined in controller/serve
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMotorsControllerClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SetJoints(ctx, &pb.Angles{Angles: tgt_angles})
	if err != nil {
		log.Fatalf("could not send: %v", err)
	}
	log.Printf("Moving to (°): %v", r.GetAngles())
}