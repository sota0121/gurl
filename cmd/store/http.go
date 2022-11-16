package store

type GurlResponse struct {
	Status    string
	Headers   map[string]string
	Body      string
	BodyBytes []byte
}

type GurlRequest struct {
	Method  string
	Url     string
	Headers map[string]string
	Body    string
}
