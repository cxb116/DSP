package utils

import "time"

func Gen10MinKeyTime(t time.Time) string {
	aligned := t.Truncate(10 * time.Minute)
	return aligned.Format("200601021504")
}
