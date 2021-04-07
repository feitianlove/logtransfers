package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Kafka *KafkaConfig
}
type KafkaConfig struct {
	Address         string
	SystemInfoTopic string
	WebTopic        string
}

func DefaultConfig() *Config {
	return &Config{
		Kafka: &KafkaConfig{
			Address:         "",
			SystemInfoTopic: "",
			WebTopic:        "",
		},
	}
}
func NewConfig(filePath string) (*Config, error) {
	cfg := DefaultConfig()
	if _, err := toml.DecodeFile(filePath, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
