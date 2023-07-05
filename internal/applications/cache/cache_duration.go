package cache

import (
	"github.com/labstack/gommon/log"
	"myapp/configs"
	"time"
)

func CachingShortPeriod() time.Duration {
	return convertStringToDuration(configs.TtlShortPeriod)
}

func CachingMediumPeriod() time.Duration {
	return convertStringToDuration(configs.TtlMediumPeriod)
}

func CachingLongPeriod() time.Duration {
	return convertStringToDuration(configs.TtlLongPeriod)
}

func convertStringToDuration(ttl string) time.Duration {
	duration, err := time.ParseDuration(ttl)
	if err != nil {
		log.Errorf("failed parsing duration", err)
		return time.Hour
	}

	return duration
}
