package s3log

import (
	"net/http"
	"time"
)

// New prepares a Entry for a new request.
// This marks the start of the request for turnaround time purposes.
func New() Entry {
	e := Entry{}
	e.Time = time.Now()
	return e
}

// Done marks the end of a request.
func (e *Entry) Done() {
	now := time.Now()
	e.Total = now.Sub(e.Time)
}

// TurnaroundDone marks the end of request processing
// and the start of the response.
func (e *Entry) TurnaroundDone() {
	now := time.Now()
	e.Turnaround = now.Sub(e.Time)
}

// SetFromRequest sets Entry values from a net/http request object.
func (e *Entry) SetFromRequest(r *http.Request) {
	e.RequestURI = r.RequestURI
}
