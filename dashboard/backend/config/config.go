package config

import (
	"github.com/spf13/viper"
)

func Init() {
	viper.AutomaticEnv()

	viper.SetDefault("CLICKHOUSE_READ_TIMEOUT", 10)
	viper.SetDefault("CLICKHOUSE_WRITE_TIMEOUT", 20)
}
