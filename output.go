package s3log

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

func quotify(s string) string {
	return "\"" + s + "\""
}

func fmtIntStr(i int64) string {
	var s string
	if i != int64(0) {
		s = strconv.FormatInt(i, 10)
	} else {
		s = "-"
	}
	return s
}

func (e *Entry) String() string {
	var s string
	s += e.Owner + " "
	s += e.Bucket + " "
	s += "[" + e.Time.Format("02/Jan/2006:15:04:05 -0700") + "]" + " "
	s += e.Remote.String() + " "
	s += e.Requester + " "
	s += e.RequestID + " "
	s += e.Operation + " "
	s += e.Key + " "
	s += quotify(e.RequestURI) + " "
	s += strconv.Itoa(e.Status) + " "
	s += e.Error + " "

	s += fmtIntStr(e.Bytes) + " "
	s += fmtIntStr(e.Size) + " "

	s += fmtIntStr(int64(e.Total/time.Millisecond)) + " "
	s += fmtIntStr(int64(e.Turnaround/time.Millisecond)) + " "

	if e.Referrer == "" {
		s += quotify("-")
	} else {
		s += quotify(e.Referrer)
	}
	s += " "

	s += quotify(e.UserAgent) + " "

	if e.Version == "" {
		s += "-"
	} else {
		s += e.Version
	}
	return s
}

func (e *Entry) Print() {
	fmt.Println(e.String())
}

func (e *Entry) Write(w io.Writer) (n int, err error) {
	return io.WriteString(w, e.String())
}
