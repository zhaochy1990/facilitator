package config

type ServerConfig struct {
	Http HttpConfig `mapstructure:"http" json:"http" yaml:"http"`
}

type HttpConfig struct {
	Port string `mapstructure:"port" json:"port" yaml:"port"`
	Mode string `mapstructure:"mode" json:"mode" yaml:"mode"`
}
