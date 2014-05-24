package s3log

import (
	"net"
	"time"
)

type Entry struct {
	Owner      string
	Bucket     string
	Time       time.Time
	Remote     net.IPAddr
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
