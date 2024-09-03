/*
Copyright © 2024 Antoine Harle antoine.harle@proton.me

*/
package controller

import (
	"fmt"

	"github.com/spf13/cobra"

	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	pb "github.com/AntoineHX/multi-motors-controller/src/proto"
)

var (
	port uint16
)
// Cobra CLI 
// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve command descritpion",
	Long: `serve command descritpion`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called with: Port ", port)
		fmt.Println(cmd.CommandPath())
		serve()
	},
}

func init() {
	controllerCmd.AddCommand(serveCmd)
	// motorCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	serveCmd.Flags().Uint16VarP(&port, "port", "p", 8080, "Port")
}

//gRPC server
// server is used to implement MotorsControllerServer.
type server struct {
	pb.UnimplementedMotorsControllerServer
}

func (s *server) SetJoints(ctx context.Context, in *pb.Angles) (*pb.Angles, error) {
	log.Printf("Received: %v", in.GetAngles())
	return &pb.Angles{Angles: in.GetAngles()}, nil
}

//TODO: Fix compiling issue with Empty message
// func (s *server) GetJoints(ctx context.Context, in *pb.emptypb.Empty) (*pb.Angles, error) {
// 	var angles = []float64{0, 0, 0}
// 	log.Printf("Sending: %v", angles)
// 	return &pb.Angles{Angles: angles}, nil
// }

func serve() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMotorsControllerServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}