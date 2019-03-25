package goserve

import (
	"errors"
	"net/http"
)

func panicRecover(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() {
			r := recover()
			if r != nil {
				switch t := r.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New("Unknown error")
				}
				WriteResponse(w, http.StatusInternalServerError, err)
			}
		}()
		h.ServeHTTP(w, r)
	})
}
