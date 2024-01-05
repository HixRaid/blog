package config

import "github.com/spf13/viper"

type Config struct {
	Server *ServerConfig `mapstructure:"server"`
	DB     *DBConfig     `mapstructure:"db"`
}

type ServerConfig struct {
	Addr string `mapstructure:"addr"`
}

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
}

func LoadConfig(fileName string) (*viper.Viper, error) {
	cfgFile := viper.New()

	cfgFile.AddConfigPath("config")
	cfgFile.SetConfigName(fileName)
	cfgFile.AutomaticEnv()

	if err := cfgFile.ReadInConfig(); err != nil {
		return nil, err
	}

	return cfgFile, nil
}

func ParseConfig(cfgFile *viper.Viper) (*Config, error) {
	cfg := &Config{}

	if err := cfgFile.Unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
