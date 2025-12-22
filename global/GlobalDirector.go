package global

import "github.com/spf13/viper"

var (
	EngineViper  *viper.Viper
	EngineConfig config.Server // 引擎配置
	EngineDB     *gorm.DB

	EngineRedis *redis.Client
	EngineKafka *sarama.KafkaVersion
)
