//Package grpc is the gRPC server implementation which runs the registry service.
package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"

	api "github.com/will-rowe/registry-microservice/pkg/api/v1"
)

// RunServer runs a gRPC service to publish the registry service.
func RunServer(ctx context.Context, serverAPI api.RegistryServiceServer, port string) error {

	// announce on the local network address
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return err
	}

	// register the registry service
	// TODO: add logging to the gRPC server by passing options to NewServer
	server := grpc.NewServer()
	api.RegisterRegistryServiceServer(server, serverAPI)

	// prepare a graceful shutdown
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go func() {

		// wait for incoming shutdown signal
		for range signalChan {
			log.Println("shut down signal received")
			server.GracefulStop()
			<-ctx.Done()
		}
	}()

	// start the gRPC server
	log.Println("starting gRPC server...")
	return server.Serve(listen)
}
