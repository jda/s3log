package s3log

import (
	"net"
	"testing"
	"time"
)

/*
79a59df900b949e55d96a1e698fbacedfd6e09d98eacf8f8d5218e7cd47ef2be mybucket [06/Feb/2014:00:00:38 +0000] 192.0.2.3 79a59df900b949e55d96a1e698fbacedfd6e09d98eacf8f8d5218e7cd47ef2be 891CE47D2EXAMPLE REST.GET.LOGGING_STATUS - "GET /mybucket?logging HTTP/1.1" 200 - 242 - 11 - "-" "S3Console/0.4" -
79a59df900b949e55d96a1e698fbacedfd6e09d98eacf8f8d5218e7cd47ef2be mybucket [06/Feb/2014:00:00:38 +0000] 192.0.2.3 79a59df900b949e55d96a1e698fbacedfd6e09d98eacf8f8d5218e7cd47ef2be A1206F460EXAMPLE REST.GET.BUCKETPOLICY - "GET /mybucket?policy HTTP/1.1" 404 NoSuchBucketPolicy 297 - 38 - "-" "S3Console/0.4" -
79a59df900b949e55d96a1e698fbacedfd6e09d98eacf8f8d5218e7cd47ef2be mybucket [06/Feb/2014:00:01:00 +0000] 192.0.2.3 79a59df900b949e55d96a1e698fbacedfd6e09d98eacf8f8d5218e7cd47ef2be 7B4A0FABBEXAMPLE REST.GET.VERSIONING - "GET /mybucket?versioning HTTP/1.1" 200 - 113 - 33 - "-" "S3Console/0.4" -
79a59df900b949e55d96a1e698fbacedfd6e09d98eacf8f8d5218e7cd47ef2be mybucket [06/Feb/2014:00:01:57 +0000] 192.0.2.3 79a59df900b949e55d96a1e698fbacedfd6e09d98eacf8f8d5218e7cd47ef2be DD6CC733AEXAMPLE REST.PUT.OBJECT s3-dg.pdf "PUT /mybucket/s3-dg.pdf HTTP/1.1" 200 - - 4406583 41754 28 "-" "S3Console/0.4" -
79a59df900b949e55d96a1e698fbacedfd6e09d98eacf8f8d5218e7cd47ef2be mybucket [06/Feb/2014:00:03:21 +0000] 192.0.2.3 79a59df900b949e55d96a1e698fbacedfd6e09d98eacf8f8d5218e7cd47ef2be BC3C074D0EXAMPLE REST.GET.VERSIONING - "GET /mybucket?versioning HTTP/1.1" 200 - 113 - 28 - "-" "S3Console/0.4" -
*/

// input cases test eache example from S3 documentation:
//   http://docs.aws.amazon.com/AmazonS3/latest/dev/LogFormat.html
func TestInputCase1(t *testing.T) {
	line := `79a59df900b949e55d96a1e698fbacedfd6e09d98eacf8f8d5218e7cd47ef2be mybucket [06/Feb/2014:00:00:38 +0000] 192.0.2.3 79a59df900b949e55d96a1e698fbacedfd6e09d98eacf8f8d5218e7cd47ef2be 3E57427F3EXAMPLE REST.GET.VERSIONING - "GET /mybucket?versioning HTTP/1.1" 200 - 113 - 7 - "-" "S3Console/0.4" -`
	owner := "79a59df900b949e55d96a1e698fbacedfd6e09d98eacf8f8d5218e7cd47ef2be"
	bucket := "mybucket"
	rtime := time.Date(2014, time.February, 6, 00, 00, 38, 0, time.UTC)
	addr := net.ParseIP("192.0.2.3")
	requester := "79a59df900b949e55d96a1e698fbacedfd6e09d98eacf8f8d5218e7cd47ef2be"
	requestID := "3E57427F3EXAMPLE"

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

	t.Log(res)
}
