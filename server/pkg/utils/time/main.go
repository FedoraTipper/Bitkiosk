package time

import (
	"strconv"
	"time"
)

func ConvertMinutesToSeconds(minutes int) int{
	return minutes * 60
}

func GetExpireTime(TTL int) *time.Time {
	duration, err := time.ParseDuration(strconv.Itoa(TTL) + "m")

	if err != nil {
		return nil
	}

	t := time.Now()
	t = t.Add(duration)

	return &t
}