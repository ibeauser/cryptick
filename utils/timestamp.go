package utils

import (
	"time"
)

func FromUnixTimestamp(ts int) time.Time {
	return time.Unix(int64(ts), 0)
}
