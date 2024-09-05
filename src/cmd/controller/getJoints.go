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
// getJointsCmd represents the getJoints command
var getJointsCmd = &cobra.Command{
	Use:   "getJoints",
	Short: "Retreive motor positions",
	Long: `Retreive the positions of the motor controlled by Motor Controller server`,
	Run: func(cmd *cobra.Command, args []string) {
		GetJoints()
	},
}

func init() {
	controllerCmd.AddCommand(getJointsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getJointsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getJointsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

//gRPC Client
func GetJoints(){
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
	r, err := c.GetJoints(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not send: %v", err)
	}
	//TODO : Use Logrus
	log.Printf("Angles (°): %v", r.GetAngles())
}