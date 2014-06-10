// Work with Amazon S3-style logs
package s3log

import (
	"net"
	"regexp"
	"time"
)

// A Entry is a structured log entry that describes a S3 request
type Entry struct {
	Owner      string
	Bucket     string
	Time       time.Time
	Remote     net.IP
	Requester  string
	RequestID  string
	Operation  string
	Key        string
	RequestURI string
	Status     int
	Error      string
	Bytes      int
	Size       int
	Total      int
	Turnaround int
	Referrer   string
	UserAgent  string
	Version    string
}

var logLine = regexp.MustCompile(`[^" ]+|("[^"]*")`)
var brackets = regexp.MustCompile(`[\[\]]`)
