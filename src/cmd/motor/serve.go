/*
Copyright © 2024 Antoine Harle antoine.harle@proton.me

*/
package motor

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	pb "github.com/AntoineHX/multi-motors-controller/src/proto"
)

var(
	stateChan chan State // State channel of the simulation
	cmdVelChan chan float64 // Command channel of the simulation
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
		init_sim()
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

//gRPC server
// server is used to implement MotorServer.
type server struct {
	pb.UnimplementedMotorServer
}

//TODO: Fix compiling issue with google.protobuf.Empty message
func (s *server) SetVelocity(ctx context.Context, in *pb.Velocity) (*pb.Empty, error) {
	log.Printf("Received: %v", in.GetVelocity())
	select{
		case cmdVelChan<-in.GetVelocity(): //Send command
			return &pb.Empty{}, nil
		default: //Drop command
			log.Printf("Failed to send command to simulation")
			return &pb.Empty{}, nil //TODO: Return error
	}
}

func (s *server) GetData(ctx context.Context, in *pb.Empty) (*pb.MotorData, error) {
	var state State
	select {
		case state = <-stateChan :
			return &pb.MotorData{Angle: state.Angle, Velocity: state.Velocity, Error: state.Error}, nil
		default:
			log.Printf("No data available")
			return &pb.MotorData{Angle: 0, Velocity: 0, Error: "No Data"}, nil
	}
}

func serve() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", curr_config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMotorServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

//Motor sim
func init_sim(){
	//Initial state
	var state State
	state.Angle = 0
	state.Velocity = 0
	state.Error = ""

	//Channels
	stateChan = make(chan State, 1) // State channel of the simulation
	cmdVelChan = make(chan float64, 1)// Command channel of the simulation

	go motor_sim(state, curr_config, 1) //Run sim at 10Hz
	//close(cmdVelChan) //Stop sim
}

func motor_sim(state_init State, config Config, sim_freq float64){
	state:=state_init 
	cmd_vel:=state.Velocity
	ok:=true
	for {
		
		//Get command velocity from channel (non-blocking)
		select{
			case cmd_vel = <- cmdVelChan:
				// check if channel is closed
				if !ok { //Stop sim
					log.Printf("Command channel closed. Stopping simulation...")
					close(stateChan)
					return
				}
				break;
			default: //No command, do nothing
		}
		
		// cmdVelChan
		//Update state
		state.Velocity = cmd_vel //Infinite acceleration
		state.Angle += state.Velocity * 1/sim_freq //Neglect computation delay for delta time
		
		//TODO: Error check

		//Send new state to channel (non-blocking)
		select{
			case stateChan <- state:
				break;
			default: //No message sent (buffer full/no receiver)	
		}

		log.Printf("Motor state: %+v", state)
		//Wait for next iteration
		time.Sleep(time.Duration(float64(time.Second)/sim_freq))
	}
}