package vars

import (
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"time"
)

type ttlPeriod struct {
	ttlShortPeriod  string
	ttlMediumPeriod string
	ttlLongPeriod   string
}

var ttl *ttlPeriod

func newTtlPeriod() *ttlPeriod {
	ttl := new(ttlPeriod)

	ttlShortPeriod := viper.GetString("cache.ttl.short")
	if ttlShortPeriod == "" {
		ttlShortPeriod = "3h"
	}
	ttl.ttlShortPeriod = ttlShortPeriod

	ttlMediumPeriod := viper.GetString("cache.ttl.medium")
	if ttlMediumPeriod == "" {
		ttlMediumPeriod = "24h"
	}
	ttl.ttlMediumPeriod = ttlMediumPeriod

	ttlLongPeriod := viper.GetString("cache.ttl.long")
	if ttlLongPeriod == "" {
		ttlLongPeriod = "72h"
	}
	ttl.ttlLongPeriod = ttlLongPeriod

	return ttl
}

func init() {
	ttl = newTtlPeriod()
}

func GetTtlShortPeriod() time.Duration {
	return convertStringToDuration(ttl.getTtlShortPeriod())
}

func GetTtlMediumPeriod() time.Duration {
	return convertStringToDuration(ttl.getTtlMediumPeriod())
}

func GetTtlLongPeriod() time.Duration {
	return convertStringToDuration(ttl.getTtlLongPeriod())
}

func (cachePeriod *ttlPeriod) getTtlShortPeriod() string {
	return cachePeriod.ttlShortPeriod
}

func (cachePeriod *ttlPeriod) getTtlMediumPeriod() string {
	return cachePeriod.ttlMediumPeriod
}

func (cachePeriod *ttlPeriod) getTtlLongPeriod() string {
	return cachePeriod.ttlLongPeriod
}

func convertStringToDuration(ttl string) time.Duration {
	duration, err := time.ParseDuration(ttl)
	if err != nil {
		log.Errorf("failed parsing duration", err)
		return time.Hour
	}

	return duration
}
