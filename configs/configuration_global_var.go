package configs

import "github.com/spf13/viper"

var (
	ApplicationName = "application.name"
	TtlShortPeriod  = "3h"
	TtlMediumPeriod = "24h"
	TtlLongPeriod   = "3d"
)

// InitGlobalVariable TODO: TBC - Need to Enhancement:
func InitGlobalVariable() {
	ApplicationName = viper.GetString("application.name")
	TtlShortPeriod = viper.GetString("cache.ttl.short-period")
	TtlMediumPeriod = viper.GetString("cache.ttl.medium-period")
	TtlLongPeriod = viper.GetString("cache.ttl.long-period")
}
