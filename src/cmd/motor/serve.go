/*
Copyright © 2024 Antoine Harle antoine.harle@proton.me

*/
package motor

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	pb "github.com/AntoineHX/multi-motors-controller/src/proto"
)

//Cobra CLI
// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve command descritpion",
	Long: `serve command descritpion`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called with ID ", cmd.Flag("id").Value)
		// fmt.Println(cmd.CommandPath())
		updateConfig()
		serve()
	},
}

func init() {
	motorCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().Uint16P("port", "p", 8080, "Port")

	//TODO: Support Motor ID flag
}

func updateConfig(){
	//Find correct ID in config file
	i := 0
	config_id := fmt.Sprintf("motors.%d.id",i)
	for viper.IsSet(config_id){
		// log.Printf("%v / %v", viper.GetInt(config_id), int(motorID))
		if viper.GetInt(config_id) == int(motorID) {
			//Extract config for this motor
			err := viper.UnmarshalKey(fmt.Sprintf("motors.%d",i), &curr_config)
			if err != nil {
				log.Fatalf("unable to decode into struct, %v", err)
				break
			}

			//Sanity check
			if curr_config.Id != motorID {
				log.Fatalf("Failed to update config. Requested ID: %d. Got: %d.", motorID, curr_config.Id)
			}
			break
		}
		//Next motor
		i++
		config_id = fmt.Sprintf("motors.%d.id",i)
	}

	log.Printf("Using config: %+v", curr_config)
}

//gRPC server
// server is used to implement MotorsControllerServer.
type server struct {
	pb.UnimplementedMotorsControllerServer
}

//TODO: Fix compiling issue with google.protobuf.Empty message
func (s *server) SetVolicty(ctx context.Context, in *pb.Velocity) (*pb.Empty, error) {
	log.Printf("Received: %v", in.GetVelocity())
	return &pb.Empty{}, nil
}

func (s *server) GetData(ctx context.Context, in *pb.Empty) (*pb.MotorData, error) {
	return &pb.MotorData{Angle: 0, Velocity: 0, Error: ""}, nil
}

func serve() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", curr_config.Port))
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