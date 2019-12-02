package transport

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/gustavolopess/locker/src/identity"
	"github.com/gustavolopess/locker/src/identity_service"
)

type IdentityEndpoints struct {
	CreateIdentity endpoint.Endpoint
	GetIdentityByID endpoint.Endpoint
	GetIdentityByFingerprint endpoint.Endpoint
}

func MakeIdentityEndpoints(s identity_service.IdentityService) IdentityEndpoints {
	return IdentityEndpoints{
		CreateIdentity:           makeCreateIdentityEndpoint(s),
		GetIdentityByID:          makeGetIdentityByIDEndpoint(s),
		GetIdentityByFingerprint: makeGetIdentityByFingerprintEndpoint(s),
	}
}

func makeCreateIdentityEndpoint(s identity_service.IdentityService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateIdentityRequest)
		newIdentity := identity.Identity{
			Fingerprint: req.Fingerprint,
		}
		err = s.CreateIdentity(newIdentity)
		return CreateIdentityResponse{err}, err
	}
}

func makeGetIdentityByIDEndpoint(s identity_service.IdentityService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetIdentityByIDRequest)
		idt, err := s.GetIdentityByID(req.ID)
		return GetIdentityByIDResponse{
			Identity: idt,
			Err:      err,
		}, err
	}
}

func makeGetIdentityByFingerprintEndpoint(s identity_service.IdentityService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetIdentityByFingerprintRequest)
		idt, err := s.GetIdentityByFingerprint(req.Fingerprint)
		return GetIdentityByFingerprintResponse{
			Identity: idt,
			Err:      err,
		}, err
	}
}