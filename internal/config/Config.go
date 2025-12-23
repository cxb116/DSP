package config

type Server struct {
	Version      string `json:"version" yaml:"version"`
	Port         string `json:"port" yaml:"port"`
	PriceEncrypt string `json:"price_encrypt" yaml:"price_encrypt"`
	WorkerSize   int    `json:"workerSize" yaml:"workerSize"`
	TaskQueue    int    `json:"taskQueue" yaml:"taskQueue"`
	Redis        Redis  `json:"redis" yaml:"redis"`
	Database     DbBase `json:"database" yaml:"database"`
	Kafka        Kafka  `json:"kafka" yaml:"kafka"`
}
