package dateutils

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
	apiDBLayout   = "2006-01-02 15:04:05"
)

//GetNow get current time in UTC
func GetNow() time.Time {
	return time.Now().UTC()
}

//GetNowString get current time in string
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

//GetNowDBFormat get current time in string for Database insertion
func GetNowDBFormat() string {
	return GetNow().Format(apiDBLayout)
}
