package utility

import "time"

const (
	dateFormat = "2006-01-02T15-04-05"
)

func FormattedNow() string {
	return time.Now().UTC().Format(dateFormat)
}
