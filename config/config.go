package config

import "github.com/zhaochy1990/x/logger"

type Config struct {
	Server       ServerConfig        `mapstructure:"server" json:"server" yaml:"server"`
	LoggerConfig logger.LoggerConfig `mapstructure:"logger" json:"logger" yaml:"logger"`
	Persister    PersisterConfig     `mapstructure:"persister" json:"persister" yaml:"persister"`
}

type PersisterConfig struct {
	Storage string `mapstructure:"storage" json:"storage" yaml:"storage"`

	Mongo PersisterMongo `mapstructure:"mongo" json:"mongo" yaml:"mongo"`
}

type PersisterMongo struct {
	URL string `mapstructure:"url" json:"url" yaml:"url"`
}
