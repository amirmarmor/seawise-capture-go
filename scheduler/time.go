package scheduler

import (
	"fmt"
	"strings"
	"time"
)

const layout = "2006-01-02T15:04:05"

func NiceTimeNow() string {
	t := time.Now()
	return NiceTime(&t)
}

func NiceTime(t *time.Time) string {
	if t == nil {
		return "nil"
	}
	return t.UTC().Format(layout)
}

func ParseTime(text string) (*time.Time, error) {
	parsed, err := time.Parse(layout, text)
	if err != nil {
		return nil, fmt.Errorf("parse time failed: %v", err)
	}
	return &parsed, nil
}

func NiceTimes(times []time.Time) string {
	var descriptions []string
	for i := 0; i < len(times); i++ {
		descriptions = append(descriptions, NiceTime(&times[i]))
	}
	return strings.Join(descriptions, ",")
}

func NiceDuration(d time.Duration) string {
	if d == 24*time.Hour {
		return "1d"
	}
	if d == 7*24*time.Hour {
		return "1w"
	}
	s := d.String()
	s = strings.Replace(s, "m0s", "m", 1)
	s = strings.Replace(s, "h0m", "h", 1)
	return s
}
