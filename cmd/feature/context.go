package feature

import (
	"log"
	"net/http"
)

// ReqContext is the http request context of the command.
type ReqContext struct {
	// method is the http request method corresponding to -X or --request.
	Method string
	// form is the http request form corresponding to -F or --form.
	form string
	// data is the data of the request corresponding to -d or --data.
	data string
	// [Not Implemented] fileToUpload is the file to upload corresponding to --upload-file.
	fileToUpload string
	// user is the user corresponding to -u or --user.
	user string
	// userAgent is the user agent corresponding to -A or --user-agent.
	userAgent string
}

// NewReqContext returns a new request context.
func NewReqContext(
	optRequest,
	optForm,
	optData,
	optUploadFile,
	optUser,
	optUserAgent string) *ReqContext {
	// Set default values
	if optRequest == "" {
		optRequest = http.MethodGet
	}

	// Check if the request method is supported.
	if !isMethodSupported(optRequest) {
		log.Fatal("Unsupported request method: ", optRequest)
		return nil
	}

	return &ReqContext{
		Method:       optRequest,
		form:         optForm,
		data:         optData,
		fileToUpload: optUploadFile,
		user:         optUser,
		userAgent:    optUserAgent,
	}
}

// isMethodSupported checks if the request method is supported.
func isMethodSupported(method string) bool {
	switch method {
	case http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodConnect, http.MethodOptions, http.MethodTrace:
		return true
	}
	return false
}
