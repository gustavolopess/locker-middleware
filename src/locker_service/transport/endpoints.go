package transport

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/gustavolopess/locker/src/locker_service"
)

// LockerEndpoints holds all Go Kit endpoints for locker service
type LockerEndpoints struct {
	CreateLocker endpoint.Endpoint
	GetLockerByID endpoint.Endpoint
		OpenLocker endpoint.Endpoint
	CloseLocker endpoint.Endpoint
}

func MakeLockerEndpoints(s locker_service.LockerService) LockerEndpoints {
	return LockerEndpoints{
		CreateLocker:  makeCreateLockerEndpoint(s),
		GetLockerByID: makeGetLockerByIDEndpoint(s),
		OpenLocker:    makeOpenLockerEndpoint(s),
		CloseLocker:   makeCloseLockerEndpoint(s),
	}
}

func makeCreateLockerEndpoint(s locker_service.LockerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateLockerRequest)
		id, err := s.CreateLocker(ctx, req.ID)
		return CreateLockerResponse{
			ID:  id,
			Err: err,
		}, nil
	}
}

func makeGetLockerByIDEndpoint(s locker_service.LockerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetLockerByIDRequest)
		loc, err := s.GetLockerByID(ctx, req.ID)
		return GetLockerByIDResponse{
			Locker: loc,
			Err:   err,
		}, nil
	}
}

func makeOpenLockerEndpoint(s locker_service.LockerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(OpenLockerRequest)
		err := s.OpenLocker(ctx, req.ID)
		return OpenLockerResponse{Err: err}, nil
	}
}

func makeCloseLockerEndpoint(s locker_service.LockerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CloseLockerRequest)
		err := s.CloseLocker(ctx, req.ID)
		return CloseLockerResponse{Err: err}, nil
	}
}



