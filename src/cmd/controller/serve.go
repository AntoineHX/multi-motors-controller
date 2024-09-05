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
	"slices"
	"math"

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

//TODO: Coroutine to regulate motor velocities
func (s *server) SetJoints(ctx context.Context, in *pb.Angles) (*pb.Angles, error) {
	//TODO: use a coroutines to prevent blocking the main thread	
	var tgt_angles = []float64{in.GetAngles()[0], in.GetAngles()[0], in.GetAngles()[0]} //TODO: Replace by in.GetAngles()
	
	//TODO: Check target angles limits

	//Compute velocities
	var(
		cmd_vel = []float64{} //Commanded velocities
		motor_pos = []float64{} //Current motor positions
		max_vels = []float64{} //Maximum velocities
		traj_times = []float64{} //Trajectory times
		s_traj_t float64 //Trajectory time (Synchronized motion)
		limit_vel = false //Limit velocity to slowest motor
		motor_errors =  []string{} //Motor errors message
		success bool = true //Success flag
	) 
	for i := range tgt_angles {
		//TODO: Use a coroutine to avoid blocking the main thread
		motor_state := getMotorState(i) //Get current motor state
		motor_pos = append(motor_pos, motor_state.Angle)
		max_vels = append(max_vels, motor_configs[i].Max_vel)
		motor_errors = append(motor_errors, "["+motor_state.Error + "]") //For display purpose
		if motor_state.Error != ""{ //If any motor is in fault, don't send command
			success = false
		}
		
		//Compute minimum trajectory times
		traj_times = append(traj_times, math.Abs(tgt_angles[i]-motor_pos[i])/max_vels[i])
	}
	log.Printf("Requested joint positions: %v -> %v", motor_pos, tgt_angles)
	// log.Printf("t: %v - %v", slices.Min(traj_times), traj_times)
	if !success{ //If any motor is in fault, stop here
		log.Printf("Motor error preventing command: %v", motor_errors)
		return &pb.Angles{Angles: motor_pos}, nil //Return motor positions TODO: Return error message
	}

	//Compute maximal velocities
	s_traj_t = slices.Min(traj_times) //Minimum trajectory time
	//TODO: Prevent NaN in case of 0 trajectory time
	for i := range tgt_angles {
		cmd_vel = append(cmd_vel, (tgt_angles[i]-motor_pos[i])/s_traj_t)
		if math.Abs(cmd_vel[i]) > max_vels[i] { // Velocity needs to be limited (Flag)
			limit_vel = true
		}
	}
	// log.Printf("cmd_vel: %v - %v", cmd_vel, limit_vel)
	if limit_vel { //Limit velocity to slowest motor
		s_traj_t = slices.Max(traj_times) //Trajectory time of slowest motor
		for i := range cmd_vel  {
			cmd_vel[i]= (tgt_angles[i]-motor_pos[i])/s_traj_t
		}
	}

	//Send command to motors
	log.Printf("Requested joints velocities (%v s): %v", s_traj_t, cmd_vel)
	for i, vel := range cmd_vel { 		
		setMotorVel(i, vel)
	}

	//Stop motors after trajectory time
	go stopMotors(time.Duration(float64(time.Second)*s_traj_t))

	return &pb.Angles{Angles: tgt_angles}, nil
}

func stopMotors(delay time.Duration) {
	timer := time.NewTimer(delay)
	<-timer.C //Block until delay is over
	for i := range motor_configs { 		
		setMotorVel(i, 0)
	}
}

func setMotorVel(idx int, vel float64){
	//TODO: Only declare client once per motor
	// Set up a connection to the server.
	var addr = fmt.Sprintf("%s:%d", ip, motor_configs[idx].Port)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMotorClient(conn)

	// Contact the server.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.SetVelocity(ctx, &pb.Velocity{Velocity: vel})
	if err != nil {
		log.Fatalf("could not send: %v", err)
	}
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
	//TODO: Only declare client once per motor
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