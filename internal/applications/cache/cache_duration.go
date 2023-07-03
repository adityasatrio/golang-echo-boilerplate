package cache

import (
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"myapp/configs"
	"time"
)

func CachingShortPeriod() time.Duration {
	return convertStringToDuration("cache.ttl.short-period")
}

func CachingMediumPeriod() time.Duration {
	return convertStringToDuration("cache.ttl.medium-period")
}

func CachingLongPeriod() time.Duration {
	return convertStringToDuration("cache.ttl.long-period")
}

func convertStringToDuration(ttl string) time.Duration {
	configs.BuildConfigEnv()
	duration, err := time.ParseDuration(viper.GetString(ttl))
	if err != nil {
		log.Errorf("failed parsing duration", err)
		return time.Hour
	}

	return duration
}
