package identity

type Identity struct {
	ID string `json:"id"`
	Fingerprint `json:"fingerprint"`
}

type Fingerprint struct {
	Payload []byte `json:"payload"`
}

