# Registry Microservice
[![tests](https://github.com/will-rowe/registry-microservice/actions/workflows/tests.yml/badge.svg)](https://github.com/will-rowe/registry-microservice/actions/workflows/tests.yml)
[![godoc](https://godoc.org/github.com/will-rowe/registry-microservice?status.svg)](https://godoc.org/github.com/will-rowe/registry-microservice)
[![goreportcard](https://goreportcard.com/badge/github.com/will-rowe/registry-microservice)](https://goreportcard.com/report/github.com/will-rowe/registry-microservice)

***

## About

This is a basic registry microservice that stores information about study participants. It includes a gRPC API that supports adding, updating, removing and retrieving participant information.

### Considerations/constraints

* unique reference numbers are allocated to participants by another microservice
* no authentication is required
* only one instance of the service is required
* no persistance between service shutdowns is required

## Implementation

This implementation is written in Go and tested using versions 1.14, 1.15 and 1.16.

### Design choices

* gRPC over REST

The main reason for selecting gRPC for the microservice API is ease of development, particularly with Go. gRPC is also performant and type safe, there is less boilerplate and it is language agnostic. One of the main drawbacks to gRPC over REST is that it has less support. If wanted, we can add REST API server using the gRPC gateway plugin which will save re-implementing the service but will lose some of the benefits of gRPC.

* data storage

As persistance isn't required the gRPC server implementation just uses a map of structs to hold participant data, where keys are the participant reference number (string) and the value is the participant data (struct). The server implementation uses a mutex to protect from concurrent RW errors. As a next step, I'd consider adding a simple, perisistant key-value store (such as [badger](https://github.com/dgraph-io/badger) or [bitcask](https://github.com/prologic/bitcask)), before using an ORM (such as [pg](https://github.com/go-pg/pg)) should an iteration on the requirements need a more fully fledged solution for persistence.

* logging

As it is currently a minimal working example, the standard libary has been used to incorporate basic logging. As a next step it would be good to implement richer logging, particularly by the gRPC server in response to incoming requests. This can be done using gRPC middleware, e.g. with [this package](https://github.com/grpc-ecosystem/go-grpc-middleware).

### Dependencies

As well as the external Go packages listed in [go.mod](./go.mod), the following tools and packages are required to build the microservice executable and documentation:

* Make
* Go toolchain
* protoc
* protoc-gen-go
* protoc-gen-doc

### Installing

For ease of development, installation is handled by the [Makefile](Makefile):

```
make all
```

This command will:
* compile the proto files for Go
* compile the gRPC API docs
* run fmt, lint and vet tools on the Go code
* run the unit tests
* build the Go executable

There is also a containerised version of the service available which is built via a [Github Action](.github/workflows/docker.yml). It can be obtained from Dockerhub:

```
docker pull willrowe/registry-microservice:latest
docker run -p 9090:9090 willrowe/registry-microservice:latest
```

### Testing

Unit tests are available for the service implementation. In addition several Go tools are used (Go lint, vet, fmt) to check the codebase. All these can be run using:

```
make test
```

To test the gRPC code without having to connect to a real server we use the [mock package](https://github.com/golang/mock); the mock class was generated using:

```
mockgen github.com/will-rowe/registry-microservice/internal/api/v1 RegistryServiceClient > internal/mock/client_mock.go
```

### Running

The client and server are available in a single binary, `registry`.

To run the server:

```
registry serve
```

To make client requests to a running server:

```
registry client -r [request] <participant-id>
```

The `-r` option supports `create`, `retrieve`, `update` and `delete`.

To retrieve a partcipant from the registry:

```
registry client -r retrieve REF-123
```

### Documentation

API documentation can be found [here](api/docs/v1/registryService.md). Implementation documentation can be found [here](https://godoc.org/github.com/will-rowe/registry-microservice).

### Limitations

* basic logging only
* no entry validation
* no HTTP/REST support