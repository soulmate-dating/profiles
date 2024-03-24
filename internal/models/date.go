package models

import "time"

const DateLayout = "2006-01-02"

func ParseDate(s string) (time.Time, error) {
	return time.Parse(DateLayout, s)
}
