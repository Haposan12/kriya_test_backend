package times

import "time"

func Now(hourOffset int, format string) string  {
	return time.Now().UTC().Add(time.Hour * time.Duration(hourOffset)).Format(format)
}