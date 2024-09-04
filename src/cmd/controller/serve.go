/*
Copyright © 2024 Antoine Harle antoine.harle@proton.me

*/
package controller

import (
	"fmt"
	"context"
	"log"
	"time"
	"net"

	motor "github.com/AntoineHX/multi-motors-controller/src/cmd/motor"

	"github.com/spf13/cobra"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/AntoineHX/multi-motors-controller/src/proto"
)

var (
	ip string = "localhost" //localhost=127.0.0.1
	port uint16 = 8080 //TODO: Share with config file or env var
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
		updateConfig()
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
	//TODO: use a coroutines to prevent blocking the main thread
	log.Printf("Received: %v", in.GetAngles())
	return &pb.Angles{Angles: in.GetAngles()}, nil
}

//TODO: Fix compiling issue with google.protobuf.Empty message
func (s *server) GetJoints(ctx context.Context, in *pb.Empty) (*pb.Angles, error) {
	//TODO: use a coroutines to prevent blocking the main thread
	var angles = []float64{}
	for i, _ := range motor_configs {
		angles = append(angles,getMotorState(i).Angle)
	}
	
	// log.Printf("Sending: %v", angles)
	return &pb.Angles{Angles: angles}, nil
}

func getMotorState(idx int)(motor.State){
	//TODO: Check if motor server is running
	//TODO: Only declare once per motor
	// Set up a connection to the server.
	var addr = fmt.Sprintf("%s:%d", ip, motor_configs[idx].Port)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMotorClient(conn)

	// Contact the server and return its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetData(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not send: %v", err)
	}
	return motor.State{Angle: r.GetAngle(), Velocity: r.GetVelocity(), Error: r.GetError()}
}

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