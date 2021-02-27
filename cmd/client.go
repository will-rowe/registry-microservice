package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"

	api "github.com/will-rowe/registry-microservice/pkg/api/v1"
)

const (
	// layoutISO is the date format for collecting the DoB from participants
	layoutISO = "2006-01-02"
)

// command line arguments
var (
	serverAddress *string // address of the server hosting the registry service
	serverRequest *string // CRUD operation to perform
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client -r [create|retrieve|update|delete] <reference_number>",
	Short: "A client to send requests to the registry server",
	Long: `Send requests to the registry server using gRPC.

	The client connects to the server and makes a CRUD request
	for participants held in the registry.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runClient(args[0])
	},
}

// init the command line arguments and add the subcommand to the root
func init() {
	serverAddress = clientCmd.Flags().StringP("serverAddress", "s", fmt.Sprintf("%s:%s", DefaultServerAddress, DefaultgRPCport), "address of the server hosting the registry service")
	serverRequest = clientCmd.Flags().StringP("request", "r", "", "server request (create|retrieve|update|delete)")
	clientCmd.MarkFlagRequired("request")
	clientCmd.MarkFlagRequired("refNum")
	rootCmd.AddCommand(clientCmd)
}

// createParticipant will collect data from a user
// and return a Participant and any error.
func createParticipant(ref string) (*api.Participant, error) {
	if len(ref) < 1 {
		return nil, errors.New("reference number is required for a participant")
	}
	p := &api.Participant{
		Id: ref,
	}

	// collect participant data from stdin and add it to the participant
	fmt.Printf("collecting information for participant (%v)\n", ref)
	phone, address, dob := "", "", ""
	fmt.Println("enter phone number:")
	fmt.Scanln(&phone)
	fmt.Println("enter address:")
	fmt.Scanln(&address)
	fmt.Println("enter date of birth (YYYY-MM-DD):")
	fmt.Scanln(&dob)
	birthdate, err := time.Parse(layoutISO, dob)
	if err != nil {
		return nil, err
	}
	p.Dob = timestamppb.New(birthdate)
	p.Phone = phone
	p.Address = address
	return p, nil
}

// runClient connects to the client and performs CRUD operation.
func runClient(refNum string) {

	// connect to the gRPC server
	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// establish the client
	client := api.NewRegistryServiceClient(conn)

	// handle request
	switch *serverRequest {
	case "create":

		// create the Participant
		p, err := createParticipant(refNum)
		if err != nil {
			log.Fatal(err)
		}

		// create the create request
		req := &api.CreateRequest{
			ApiVersion:  DefaultAPIVersion,
			Participant: p,
		}

		// setup context
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// send request and check response
		res, err := client.Create(ctx, req)
		if err != nil {
			log.Fatalf("create request failed: %v", err)
		}
		if res.Created != true {
			log.Fatal("create request failed")
		}
		log.Printf("create request successful for: %v", refNum)
	case "retrieve":

		// create the retrieve request
		req := &api.RetrieveRequest{
			ApiVersion: DefaultAPIVersion,
			Id:         refNum,
		}

		// setup context
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// send the request and check response
		res, err := client.Retrieve(ctx, req)
		if err != nil {
			log.Fatalf("retrieve request failed: %v", err)
		}

		// print the retrieved data to STDOUT
		fmt.Fprintln(os.Stdout, res.Participant.String())
		log.Printf("retrieve request successful for: %v", refNum)
	case "update":

		// create the Participant
		p, err := createParticipant(refNum)
		if err != nil {
			log.Fatal(err)
		}

		// create the update request
		req := &api.UpdateRequest{
			ApiVersion:  DefaultAPIVersion,
			Participant: p,
		}

		// setup context
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// send request and check the response
		res, err := client.Update(ctx, req)
		if err != nil {
			log.Fatalf("update request failed: %v", err)
		}
		if res.Updated != true {
			log.Fatal("update request failed")
		}
		log.Printf("update request successful for: %v", refNum)
	case "delete":

		// create the delete request
		req := &api.DeleteRequest{
			ApiVersion: DefaultAPIVersion,
			Id:         refNum,
		}

		// setup context
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// send the request and check response
		res, err := client.Delete(ctx, req)
		if err != nil {
			log.Fatalf("delete request failed: %v", err)
		}
		if res.Deleted != true {
			log.Fatal("delete request failed")
		}
		log.Printf("delete request successful for: %v", refNum)
	default:
		log.Fatal("only create|retrieve|update|delete requests are supported")
	}
}
