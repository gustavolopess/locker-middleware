package transport

import (
	"github.com/gustavolopess/locker/src/locker"
)

type CreateLockerRequest struct {
	ID string `json:"id"`
}

type CreateLockerResponse struct {
	ID string `json:"id,omitempty"`
	Err error `json:"error,omitempty"`
}

type GetLockerByIDRequest struct {
	ID string
}

type GetLockerByIDResponse struct {
	Locker locker.Locker `json:"locker"`
	Err error `json:"error,omitempty"`
}

type OpenLockerRequest struct {
	ID string
}

type OpenLockerResponse struct {
	Err error `json:"error,omitempty"`
}

type CloseLockerRequest struct {
	ID string
}

type CloseLockerResponse struct {
	Err error `json:"error,omitempty"`
}

