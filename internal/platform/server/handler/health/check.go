package health

import (
	"net/http"
)

// CheckHandler returns an HTTP handler to perform health checks.
func CheckHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("everything is ok!"))
	}
}
