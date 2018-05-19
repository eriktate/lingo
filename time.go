package lingo

import (
	"strconv"
	"time"
)

const formatString = "2006-01-02T15:04:05"

// Time is a type alias for time.Time that handles proper JSON marshaling and unmarshaling
// for the Linode API.
type Time struct {
	time.Time
}

// MarshalJSON implements the json.Marshaler interface for the custom
// Time type.
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(t.Format(formatString))), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for the custom
// Time type.
func (t *Time) UnmarshalJSON(data []byte) error {
	dateString, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}

	newTime, err := time.Parse(formatString, dateString)
	if err != nil {
		return err
	}

	t.Time = newTime
	return nil
}
