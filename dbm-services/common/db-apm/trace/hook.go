package trace

import (
	"context"
	"fmt"

	log "db-apm/log"
	"github.com/spf13/viper"
)

// setDefaultConfig
func setDefaultConfig() {
	viper.SetDefault(KeysConfigPath, nil)
	viper.SetDefault(OtlpHostConfigPath, "127.0.0.1")
	viper.SetDefault(OtlpPortConfigPath, "4317")
	viper.SetDefault(OtlpTokenConfigPath, "")
	viper.SetDefault(OtlpTypeConfigPath, "grpc")

	viper.SetDefault(ServiceNameConfigPath, "drs")
	viper.SetDefault(DataIDConfigPath, 0)
}

// InitConfig
func InitConfig() {

	DataID = viper.GetInt64(DataIDConfigPath)

	for key, value := range configLabels {
		log.Debugf(context.TODO(), "key->[%s] value->[%s] now is added to labels", key, value)
		labels[key] = value
	}

	otlpHost = viper.GetString(OtlpHostConfigPath)
	otlpPort = viper.GetString(OtlpPortConfigPath)
	otlpToken = viper.GetString(OtlpTokenConfigPath)
	log.Infof(context.TODO(), "trace will Otlp to host->[%s] port->[%s] token->[%s]", otlpHost, otlpPort, otlpToken)

	OtlpType = viper.GetString(OtlpTypeConfigPath)
	log.Infof(context.TODO(), "trace will Otlp as %s type", OtlpType)

	ServiceName = viper.GetString(ServiceNameConfigPath)
	log.Infof(context.TODO(), "trace will Otlp service name:%s", ServiceName)
}

// init
func init() {
	fmt.Println("trace hook init.")
	setDefaultConfig()
	InitConfig()
}
