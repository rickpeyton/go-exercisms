package gigasecond

import (
	"time"
)

const (
	TestVersion = 2
)

var Birthday = time.Date(1980, 11, 9, 0, 0, 0, 0, time.UTC)

// Add a gigasecond to the time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1000000000 * time.Second)
}
