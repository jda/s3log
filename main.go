// Work with Amazon S3-style logs
package s3log

import (
	"net"
	"regexp"
	"time"
)

// A Entry is a structured log entry that describes a S3 request
type Entry struct {
	Owner      string    // user ID of bucket owner
	Bucket     string    // bucket name
	Time       time.Time // time when request was recieved
	Remote     net.IP    // IP address of requester
	Requester  string    // user ID of requester
	RequestID  string    // request ID
	Operation  string
	Key        string // key requested from bucket
	RequestURI string
	Status     int           // HTTP status code
	Error      string        // Error code
	Bytes      int64         // Bytes sent to requester
	Size       int64         // size of object requested
	Total      time.Duration // time spent serving request
	Turnaround time.Duration // time spent handling request before response is sent
	Referrer   string        // HTTP referred
	UserAgent  string        // HTTP UserAgent
	Version    string        // Request version ID
}

var logLine = regexp.MustCompile(`[^" ]+|("[^"]*")`)
var brackets = regexp.MustCompile(`[\[\]]`)
