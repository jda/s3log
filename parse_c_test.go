package s3log

import (
	"net"
	"testing"
	"time"
)

// input cases test eache example from S3 documentation:
//   http://docs.aws.amazon.com/AmazonS3/latest/dev/LogFormat.html
func TestInputCaseC(t *testing.T) {
	line := `79a59df900b949e55d96a1e698fbacedfd6e09d98eacf8f8d5218e7cd47ef2be mybucket [06/Feb/2014:00:00:38 +0000] 192.0.2.3 79a59df900b949e55d96a1e698fbacedfd6e09d98eacf8f8d5218e7cd47ef2be A1206F460EXAMPLE REST.GET.BUCKETPOLICY - "GET /mybucket?policy HTTP/1.1" 404 NoSuchBucketPolicy 297 - 38 - "-" "S3Console/0.4" -`
	owner := "79a59df900b949e55d96a1e698fbacedfd6e09d98eacf8f8d5218e7cd47ef2be"
	bucket := "mybucket"
	rtime := time.Date(2014, time.February, 6, 00, 00, 38, 0, time.UTC)
	addr := net.ParseIP("192.0.2.3")
	requester := "79a59df900b949e55d96a1e698fbacedfd6e09d98eacf8f8d5218e7cd47ef2be"
	requestID := "A1206F460EXAMPLE"
	operation := "REST.GET.BUCKETPOLICY"
	key := "-"
	ruri := "GET /mybucket?policy HTTP/1.1"
	hstatus := 404
	errorCode := "NoSuchBucketPolicy"
	bytesSent := int64(297)
	objectSize := int64(0)
	totalTime := time.Duration(38) * time.Millisecond
	turnaround := time.Duration(0) * time.Millisecond
	referrer := ""
	ua := "S3Console/0.4"
	version := ""

	res, err := Parse(line)
	if err != nil {
		t.Errorf("could not parse line because %s", err)
	}

	if res.Owner != owner {
		t.Errorf("Owner: got %s. should be %s", res.Owner, owner)
	}

	if res.Bucket != bucket {
		t.Errorf("Bucket: got %s. should be %s", res.Bucket, bucket)
	}

	if res.Time != rtime {
		t.Errorf("Time: got %s. should be %s", res.Time, rtime)
	}

	if res.Remote.String() != addr.String() {
		t.Errorf("Remote: got %s. should be %s", res.Remote, addr)
	}

	if res.Requester != requester {
		t.Errorf("Requester: got %s. should be %s", res.Requester, requester)
	}

	if res.RequestID != requestID {
		t.Errorf("Request ID: got %s. should be %s", res.RequestID, requestID)
	}

	if res.Operation != operation {
		t.Errorf("Operation: got %s. should be %s", res.Operation, operation)
	}

	if res.Key != key {
		t.Errorf("Key: got %s. should be %s", res.Key, key)
	}

	if res.RequestURI != ruri {
		t.Errorf("Request URI: got %s. should be %s", res.RequestURI, ruri)
	}

	if res.Status != hstatus {
		t.Errorf("HTTP Status: got %d. should be %d", res.Status, hstatus)
	}

	if res.Error != errorCode {
		t.Errorf("Error code: got %s. should be %s", res.Error, errorCode)
	}

	if res.Bytes != bytesSent {
		t.Errorf("Bytes sent: got %d. should be %d", res.Bytes, bytesSent)
	}

	if res.Size != objectSize {
		t.Errorf("Object size: get %d. should be %d", res.Size, objectSize)
	}

	if res.Total != totalTime {
		t.Errorf("Total time: got %d. should be %d", res.Total, totalTime)
	}

	if res.Turnaround != turnaround {
		t.Errorf("Turnaround time: got %d. should be %d", res.Turnaround, turnaround)
	}

	if res.Referrer != referrer {
		t.Errorf("Referrer: got %s. should be %s", res.Referrer, referrer)
	}

	if res.UserAgent != ua {
		t.Errorf("UserAgent: got %s. should be %s", res.UserAgent, ua)
	}

	if res.Version != version {
		t.Errorf("Version: got %s. should be %s", res.Version, version)
	}

	rOut := res.String()
	if rOut != line {
		t.Errorf("Output line does not match input:\n In: %s\nOut: %s\n", line, rOut)
	}
}
