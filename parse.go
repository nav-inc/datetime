package datetime

import (
	"bytes"
	"time"
)

// ParseTime takes a string with a ISO 8601 timestamp in it and returns a time.Time.
func ParseTime(s string) (time.Time, error) {
	p := newParser(bytes.NewBuffer([]byte(s)))
	return p.parse()
}