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
// listenCmd represents the listen command
var listenCmd = &cobra.Command{
	Use:   "listen",
	Short: "listen command descritpion",
	Long: `listen command descritpion`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listen called with ID: ", cmd.Flag("id").Value)
		updateConfig()
		listen() //Blocking call
	},
}

func init() {
	motorCmd.AddCommand(listenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listenCmd.Flags().Uint16("id", 0, "Identifier number")
}

// Could have been done with a coroutine and a range on a channel if launched from the server thread
func listen(){
	// Set up a connection to the server.
	var addr = fmt.Sprintf("%s:%d", ip, curr_config.Port) //Defined in motor
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMotorClient(conn)

	for{
		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.GetData(ctx, &pb.Empty{})
		if err != nil {
			log.Fatalf("could not send: %v", err)
		}

		//TODO : Use Logrus
		log.Printf("Motor state: %f° | %f°/s | %s", r.GetAngle(), r.GetVelocity(), r.GetError())

		// Wait for a second
		time.Sleep(time.Second)
	}
}