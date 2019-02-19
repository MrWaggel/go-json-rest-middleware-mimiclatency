// Package mimiclatency provides a middleware that produces latency for the
// http://ant0ine.github.io/go-json-rest/ package.
package mimiclatency

import (
	"github.com/ant0ine/go-json-rest/rest"
	"time"
)

// MimicLatencyMiddleWare mimics latency on the server. This is ideally not
// suited for production use.
type MimicLatencyMiddleWare struct {

	// Default value of duration is 0 seconds
	Duration time.Duration
}

// MiddlewareFunc makes MimicLatencyMiddleWare implement the Middleware interface.
func (m MimicLatencyMiddleWare) MiddlewareFunc(h rest.HandlerFunc) rest.HandlerFunc {
	return func(w rest.ResponseWriter, r *rest.Request) {
		time.Sleep(m.Duration)
		h(w, r)
	}
}
