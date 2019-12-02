package transport

import "github.com/gustavolopess/locker/src/identity"

type CreateIdentityRequest struct {
	Fingerprint string `json:"fingerprint"`
}

type CreateIdentityResponse struct {
	Err error `json:"error"`
}

type GetIdentityByIDRequest struct {
	ID string `json:"id"`
}

type GetIdentityByIDResponse struct {
	Identity identity.Identity `json:"identity"`
	Err error `json:"error"`
}

type GetIdentityByFingerprintRequest struct {
	Fingerprint string `json:"fingerprint"`
}

type GetIdentityByFingerprintResponse struct {
	Identity identity.Identity `json:"identity"`
	Err error `json:"error"`
}
