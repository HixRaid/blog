package config

import "github.com/spf13/viper"

type Config struct {
	Server *ServerConfig
	DB     *DBConfig
}

type ServerConfig struct {
	Addr string `yaml:"addr"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
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
