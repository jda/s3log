package s3log

import (
	"fmt"
	//"net"
	"testing"
	"time"
)

// input cases test eache example from S3 documentation:
//   http://docs.aws.amazon.com/AmazonS3/latest/dev/LogFormat.html
func TestRequestCaseA(t *testing.T) {
	/*
		line := `79a59df900b949e55d96a1e698fbacedfd6e09d98eacf8f8d5218e7cd47ef2be mybucket [06/Feb/2014:00:00:38 +0000] 192.0.2.3 79a59df900b949e55d96a1e698fbacedfd6e09d98eacf8f8d5218e7cd47ef2be 3E57427F3EXAMPLE REST.GET.VERSIONING - "GET /mybucket?versioning HTTP/1.1" 200 - 113 - 7 - "-" "S3Console/0.4" -`
		owner := "79a59df900b949e55d96a1e698fbacedfd6e09d98eacf8f8d5218e7cd47ef2be"
		bucket := "mybucket"
		rtime := time.Date(2014, time.February, 6, 00, 00, 38, 0, time.UTC)
		addr := net.ParseIP("192.0.2.3")
		requester := "79a59df900b949e55d96a1e698fbacedfd6e09d98eacf8f8d5218e7cd47ef2be"
		requestID := "3E57427F3EXAMPLE"
		operation := "REST.GET.VERSIONING"
		key := "-"
		ruri := "GET /mybucket?versioning HTTP/1.1"
		hstatus := 200
		errorCode := "-"
		bytesSent := 113
		objectSize := 0
		totalTime := 7
		turnaround := 0
		referrer := ""
		ua := "S3Console/0.4"
		version := ""
	*/

	e := New()
	time.Sleep(3 * time.Second)
	e.TurnaroundDone()
	time.Sleep(1 * time.Second)
	e.Done()
	fmt.Printf("e: %+v", e)
}
