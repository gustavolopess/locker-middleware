package transport

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/gustavolopess/locker/src/locker_service"
)

// Endpoints holds all Go Kit endpoints for locker service
type Endpoints struct {
	CreateLocker endpoint.Endpoint
	GetLockerByID endpoint.Endpoint
	OpenLocker endpoint.Endpoint
	CloseLocker endpoint.Endpoint
}

func MakeEndpoints(s locker_service.Service) Endpoints {
	return Endpoints{
		CreateLocker:  makeCreateLockerEndpoint(s),
		GetLockerByID: makeGetLockerByIDEndpoint(s),
		OpenLocker:    makeOpenLockerEndpoint(s),
		CloseLocker:   makeCloseLockerEndpoint(s),
	}
}

func makeCreateLockerEndpoint(s locker_service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createLockerRequest)
		id, err := s.CreateLocker(ctx, req.Locker)
		return createLockerResponse{
			ID:  id,
			Err: err,
		}, nil
	}
}

func makeGetLockerByIDEndpoint(s locker_service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getLockerByIDRequest)
		loc, err := s.GetLockerByID(ctx, req.ID)
		return getLockerByIDResponse{
			Locker: loc,
			Err:   err,
		}, nil
	}
}

func makeOpenLockerEndpoint(s locker_service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(openLockerRequest)
		err := s.OpenLocker(ctx, req.ID)
		return openLockerResponse{Err: err}, nil
	}
}

func makeCloseLockerEndpoint(s locker_service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(closeLockerRequest)
		err := s.CloseLocker(ctx, req.ID)
		return closeLockerResponse{Err: err}, nil
	}
}



