//Package service implements the registry service API.
package service

import (
	"context"
	"sync"

	api "github.com/will-rowe/registry-microservice/internal/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	apiVersion = "1"
)

// registryService is an implementation
// of the v1.RegistryServiceServer.
type registryService struct {

	// version of API implemented by the server
	version string

	// db is the in-memory db to store participants
	db map[string]api.Participant

	// db lock
	sync.RWMutex
}

// NewRegistryService creates the registry service.
func NewRegistryService() api.RegistryServiceServer {
	return &registryService{
		version: apiVersion,
		db:      make(map[string]api.Participant),
	}
}

// checkAPI checks if requested API version is supported
// by the server.
func (rs *registryService) checkAPI(requestedAPI string) error {
	if rs.version != requestedAPI {
		return status.Errorf(codes.Unimplemented,
			"unsupported API version requested: current service implements version '%s', but version '%s' was requested", rs.version, requestedAPI)
	}
	return nil
}

// Create will create a new participant in the registry.
func (rs *registryService) Create(ctx context.Context, request *api.CreateRequest) (*api.CreateResponse, error) {

	// check we have received a supported API request
	if err := rs.checkAPI(request.GetApiVersion()); err != nil {
		return nil, err
	}

	// lock the db for RW access
	rs.Lock()
	defer rs.Unlock()

	// check if entry already exists for provided reference number
	if _, ok := rs.db[request.GetParticipant().GetId()]; ok {
		return nil, status.Errorf(codes.AlreadyExists,
			"reference number in use: participant already exists in the registry for %v", request.GetParticipant().GetId())
	}

	// TODO: validate the provided participant details

	// add the participant as an entry in the registry db
	rs.db[request.GetParticipant().GetId()] = *request.GetParticipant()

	// create a response and return
	return &api.CreateResponse{
		ApiVersion: rs.version,
		Created:    true,
	}, nil
}

// Retrieve will retrieve a participant from the registry.
func (rs *registryService) Retrieve(ctx context.Context, request *api.RetrieveRequest) (*api.RetrieveResponse, error) {

	// check we have received a supported API request
	if err := rs.checkAPI(request.GetApiVersion()); err != nil {
		return nil, err
	}

	// check if entry already exists for provided reference number
	participant, ok := rs.db[request.GetId()]
	if !ok {
		return nil, status.Errorf(codes.NotFound,
			"reference number not found: no participant entry exists in the registry for %v", request.GetId())
	}

	// create a response and return
	return &api.RetrieveResponse{
		ApiVersion:  rs.version,
		Participant: &participant,
	}, nil
}

// Update will update a participant in the registry.
// NOTE: this will update all fields, effectively calling delete and then create
// TODO: implement individual field updates for a participant
func (rs *registryService) Update(ctx context.Context, request *api.UpdateRequest) (*api.UpdateResponse, error) {

	// check we have received a supported API request
	if err := rs.checkAPI(request.GetApiVersion()); err != nil {
		return nil, err
	}

	// lock the db for RW access
	rs.Lock()
	defer rs.Unlock()

	// check if entry already exists for provided reference number
	if _, ok := rs.db[request.GetParticipant().GetId()]; !ok {
		return nil, status.Errorf(codes.NotFound,
			"reference number not found: no participant entry exists in the registry for %v", request.GetParticipant().GetId())
	}

	// delete the entry from the registry db
	delete(rs.db, request.GetParticipant().GetId())

	// TODO: validate the provided participant details

	// add the participant as an entry in the registry db
	rs.db[request.GetParticipant().GetId()] = *request.GetParticipant()

	// create a response and return
	return &api.UpdateResponse{
		ApiVersion: rs.version,
		Updated:    true,
	}, nil
}

// Delete will delete a participant from the registry.
func (rs *registryService) Delete(ctx context.Context, request *api.DeleteRequest) (*api.DeleteResponse, error) {

	// check we have received a supported API request
	if err := rs.checkAPI(request.GetApiVersion()); err != nil {
		return nil, err
	}

	// lock the db for RW access
	rs.Lock()
	defer rs.Unlock()

	// check if entry already exists for provided reference number
	if _, ok := rs.db[request.GetId()]; !ok {
		return nil, status.Errorf(codes.NotFound,
			"reference number not found: no participant entry exists in the registry for %v", request.GetId())
	}

	// delete the entry from the registry db
	delete(rs.db, request.GetId())

	// create a response and return
	return &api.DeleteResponse{
		ApiVersion: rs.version,
		Deleted:    true,
	}, nil
}
