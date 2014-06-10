package s3log

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

// Parse parses a Amazon S3-style log line into a Entry
func Parse(line string) (Entry, error) {
	e := Entry{}
	l := logLine.FindAllString(line, -1)

	e.Owner = l[0]
	e.Bucket = l[1]
	pt, err := parseTime(l[2] + " " + l[3])
	if err != nil {
		return e, err
	}
	e.Time = pt
	e.Remote = net.ParseIP(l[4])
	e.Requester = l[5]
	e.RequestID = l[6]
	e.Operation = l[7]
	e.Key = l[8]

	prequestURI := l[9]
	e.RequestURI = strings.Replace(prequestURI, `"`, "", -1)

	pstatus, err := strconv.Atoi(l[10])
	if err != nil {
		return e, err
	}
	e.Status = pstatus

	perrorRaw := l[11]
	if perrorRaw == "-" {
		e.Error = ""
	} else {
		e.Error = perrorRaw
	}
	e.Error = perrorRaw

	pbytesRaw := l[12]
	var pbytes int64
	if pbytesRaw != "-" {
		pbytes, err = strconv.ParseInt(pbytesRaw, 10, 64)
		if err != nil {
			return e, err
		}
	}
	e.Bytes = pbytes

	psizeRaw := l[13]
	var psize int64
	if psizeRaw != "-" {
		psize, err = strconv.ParseInt(l[13], 10, 64)
		if err != nil {
			return e, err
		}
	}
	e.Size = psize

	ptotalRaw := l[14]
	ptotal := 0
	if ptotalRaw != "-" {
		ptotal, err = strconv.Atoi(l[14])
		if err != nil {
			return e, err
		}
	}
	e.Total = time.Duration(ptotal) * time.Millisecond

	pturnaroundRaw := l[15]
	pturnaround := 0
	if pturnaroundRaw != "-" {
		pturnaround, err = strconv.Atoi(l[15])
		if err != nil {
			log.Panic(err)

			return e, err
		}
	}
	e.Turnaround = time.Duration(pturnaround) * time.Millisecond

	pReferrer := l[16]
	pReferrer = strings.Replace(pReferrer, `"`, "", -1)
	if pReferrer == "-" {
		pReferrer = ""
	}
	e.Referrer = pReferrer

	pUserAgent := l[17]
	e.UserAgent = strings.Replace(pUserAgent, `"`, "", -1)

	pVersion := l[18]
	if pVersion == "-" {
		pVersion = ""
	}
	e.Version = pVersion

	return e, nil
}

func parseTime(tl string) (time.Time, error) {
	tl = brackets.ReplaceAllString(tl, "")
	t, err := time.Parse("02/Jan/2006:15:04:05 -0700", tl)
	if err != nil {
		return time.Now(), fmt.Errorf("Error parsing time: %s", err)
	}

	t = t.In(time.UTC)

	return t, nil
}
