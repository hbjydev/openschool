package config

import (
	"errors"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type BusReconnectConfig struct {
	Interval   time.Duration
	MaxAttempt int
}

type BusConfig struct {
	Username             string             `yaml:"username"`
	Password             string             `yaml:"password"`
	Host                 string             `yaml:"host"`
	Port                 int                `yaml:"port"`
	Vhost                string             `yaml:"vhost,omitempty"`
	ConnectionName       string             `yaml:"connectionName,omitempty"`
	ChannelNotifyTimeout time.Duration      `yaml:"notifyTimeout,omitempty"`
	Reconnect            BusReconnectConfig `yaml:"reconnect"`
}

type Config struct {
	Addr string    `yaml:"addr"`
	Bus  BusConfig `yaml:"bus"`
}

func NewDevelopment() *Config {
	busConf := BusConfig{
		Username: "guest",
		Password: "guest",
		Host:     "127.0.0.1",
		Port:     5672,
		Vhost:    "",
		Reconnect: struct {
			Interval   time.Duration
			MaxAttempt int
		}{},
	}

	conf := Config{
		Addr: "127.0.0.1:8001",
		Bus:  busConf,
	}

	return &conf
}

func ReadConfigFile(path string) (*Config, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if stat.Mode().IsDir() {
		return nil, errors.New("config file path is a directory")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
