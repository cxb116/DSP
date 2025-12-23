package config

import "time"

type Kafka struct {
	Brokers       []string      `yaml:"brokers"`
	Frequency     time.Duration `yaml:"frequency"`
	Messages      int           `yaml:"messages"`
	ReturnSuccess bool          `yaml:"returnSuccess"`
	TopicSsp      string        `yaml:"topic_ssp"`
}
