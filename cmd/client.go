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

// command line arguments
var (
	serverAddress *string // address of the server hosting the registry service
	serverRequest *string // CRUD operation to perform
	refNum        *string // reference number for the participant
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "A client to send requests to the registry server",
	Long: `Send requests to the registry server using gRPC.

	The client connects to the server and makes a CRUD request
	for participants held in the registry.`,
	Run: func(cmd *cobra.Command, args []string) {
		runClient()
	},
}

// init the command line arguments and add the subcommand to the root
func init() {
	serverAddress = clientCmd.Flags().StringP("serverAddress", "s", fmt.Sprintf("%s:%s", DefaultServerAddress, DefaultgRPCport), "address of the server hosting the registry service")
	serverRequest = clientCmd.Flags().StringP("request", "r", "", "server request (create|retrieve|update|delete)")
	refNum = clientCmd.Flags().StringP("refnum", "i", "", "reference number for participant")
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
	fmt.Printf("collecting information for participant (%v)", ref)
	phone, address := "", ""
	var day, month, year int
	fmt.Println("enter phone number:")
	fmt.Scanln(&phone)
	fmt.Println("enter address:")
	fmt.Scanln(&address)
	fmt.Println("enter birthdate day (DD)")
	if _, err := fmt.Scanf("%d", &day); err != nil {
		return nil, fmt.Errorf("could not collect day: %v", err)
	}
	if day == 0 || day > 32 {
		return nil, errors.New("birthdate day invalid")
	}
	fmt.Println("enter birthdate month (MM)")
	if _, err := fmt.Scanf("%d", &month); err != nil {
		return nil, fmt.Errorf("could not collect day: %v", err)
	}
	if month == 0 || month > 12 {
		return nil, errors.New("birthdate month invalid")
	}
	fmt.Println("enter birthdate year (YYYY)")
	if _, err := fmt.Scanf("%d", &year); err != nil {
		return nil, fmt.Errorf("could not collect day: %v", err)
	}
	if year < 1900 || year > 2021 {
		return nil, errors.New("birthdate year invalid")
	}
	birthdate, err := time.Parse("2021-01-17", fmt.Sprintf("%d-%d-%d", year, month, day))
	if err != nil {
		return nil, err
	}
	p.Dob = timestamppb.New(birthdate)
	p.Phone = phone
	p.Address = address
	return p, nil
}

// runClient connects to the client and performs CRUD operation.
func runClient() {

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
		p, err := createParticipant(*refNum)
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
		log.Printf("create request successful for: %v", *refNum)
	case "retrieve":

		// create the retrieve request
		req := &api.RetrieveRequest{
			ApiVersion: DefaultAPIVersion,
			Id:         *refNum,
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
		log.Printf("retrieve request successful for: %v", *refNum)
	case "update":

		// create the Participant
		p, err := createParticipant(*refNum)
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
		log.Printf("update request successful for: %v", *refNum)
	case "delete":

		// create the delete request
		req := &api.DeleteRequest{
			ApiVersion: DefaultAPIVersion,
			Id:         *refNum,
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
		log.Printf("delete request successful for: %v", *refNum)
	default:
		log.Fatal("only create|retrieve|update|delete requests are supported")
	}
}
