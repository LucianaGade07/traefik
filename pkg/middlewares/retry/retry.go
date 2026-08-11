package retry

import (
	"context"
	"net"
	"net/http"
	"github.com/traefik/traefik/v2/pkg/log"
)

func (r *Retry) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	var lastErr error
	for i := 0; i <= r.attempts; i++ {
		// Check if context is canceled before retrying
		select {
		case <-req.Context().Done():
			return
		default:
		}

		// Execute the request
		r.next.ServeHTTP(rw, req)

		// Logic to check if the last attempt resulted in a connection error
		// This assumes the round-tripper or proxy sets an error state we can inspect
		if err := getTransportError(req); err != nil {
			lastErr = err
			log.FromContext(req.Context()).Debugf("Retry attempt %d failed due to connection error: %v", i+1, err)
			continue
		}
		return
	}
}

func getTransportError(req *http.Request) error {
	// Implementation to extract net.Error from the request context or proxy state
	// This is a placeholder for the actual error extraction logic
	return nil
}