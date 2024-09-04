/*
Copyright © 2024 Antoine Harle antoine.harle@proton.me

*/
package controller

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
// setJointsCmd represents the setJoints command
var setJointsCmd = &cobra.Command{
	Use:   "setJoints",
	Short: "setJoints command descritpion",
	Long: `setJoints command descritpion`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("setJoints called with target :", cmd.Flag("j").Value)
		tgt_angles, err:=cmd.Flags().GetFloat64Slice("j")
		if err!= nil {
			log.Fatalf("Failed read requested velocities %v", err)
		}else{
			SetJoints(tgt_angles)
		}
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
	setJointsCmd.Flags().Float64Slice("j", nil, "Target joint values") //or IntSlice ?
	setJointsCmd.MarkFlagRequired("j")
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
	log.Printf("Received: %v", r.GetAngles())
}