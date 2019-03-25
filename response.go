package goserve

import (
	"encoding/json"
	"net/http"
)

// ResponseBody holds information for responding to a request.  Generally
// this struct should be used for successful requests (code < 300).
type ResponseBody struct {
	Status  string      `json:"status,omitempty"`
	Message interface{} `json:"message,omitempty"`
}

// ResponseError holds information for responding to a request with an
// error (code >= 300).
type ResponseError struct {
	Code   int    `json:"code,omitempty"`
	Status string `json:"status,omitempty"`
	Error  string `json:"error,omitempty"`
}

// WriteResponse is a wrapper for writing the response to an HTTP request. The
// body of the response will change depending on if the body parameter is a string,
// an error, or a struct. If body is a string or an error, the response body will be
//			{
//				"status": <status_code_text>,
//				"message": <string/error>
//			}
// If body is a struct, it will be encoded as is.
func WriteResponse(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)

	if body != nil {
		var responseBody interface{}
		switch val := body.(type) {
		case string:
			responseBody = ResponseBody{Status: http.StatusText(status), Message: val}
		case error:
			responseBody = ResponseBody{Status: http.StatusText(status), Message: val.Error()}
		default:
			responseBody = body
		}

		if err := json.NewEncoder(w).Encode(responseBody); err != nil {
			panic(err)
		}
	}
}
