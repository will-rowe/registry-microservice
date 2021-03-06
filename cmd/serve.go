package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	server "github.com/will-rowe/registry-microservice/pkg/protocol/grpc"
	service "github.com/will-rowe/registry-microservice/pkg/service/v1"
)

// command line arguments
var (
	grpcPort *string // TCP port to listen to by the gRPC server
	logFile  *string // the log file
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the registry server",
	Long: `Run the registry server using gRPC.

Clients can then connect to the server and make CRUD requests
for participants held in the registry.`,
	Run: func(cmd *cobra.Command, args []string) {
		runServer()
	},
}

// init the command line arguments and add the subcommand to the root
func init() {
	grpcPort = serveCmd.Flags().StringP("grpcPort", "g", DefaultgRPCport, "TCP port to listen to by the gRPC server")
	logFile = serveCmd.Flags().StringP("logFile", "l", DefaultLogFile, "the file to write the server log to (use -l STDOUT for logging to standard out)")
	rootCmd.AddCommand(serveCmd)
}

// runServer sets up and runs the gRPC server and HTTP gateway
func runServer() {

	// set up the log
	if *logFile != "STDOUT" {
		file, err := os.OpenFile(*logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		log.SetOutput(file)
	}
	log.Println("registry microservice launched")

	// get top level context
	ctx := context.Background()

	// get the server API
	serverAPI := service.NewRegistryService()

	// run the server until shutdown signal received
	if err := server.RunServer(ctx, serverAPI, *grpcPort); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	log.Println("finished")
}
