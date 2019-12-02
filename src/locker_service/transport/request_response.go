package transport

import "github.com/gustavolopess/locker/src/locker"

type createLockerRequest struct {
	Locker locker.Locker
}

type createLockerResponse struct {
	ID string `json:"id,omitempty"`
	Err error `json:"error,omitempty"`
}

type getLockerByIDRequest struct {
	ID string
}

type getLockerByIDResponse struct {
	Locker locker.Locker `json:"locker"`
	Err error `json:"error,omitempty"`
}

type openLockerRequest struct {
	ID string
}

type openLockerResponse struct {
	Err error `json:"error,omitempty"`
}

type closeLockerRequest struct {
	ID string
}

type closeLockerResponse struct {
	Err error `json:"error,omitempty"`
}

