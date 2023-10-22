package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Logger *zap.Logger
	Port   string `mapstructure:"SERVER_PORT"`
	Host   string `mapstructure:"SERVER_HOST"`
}

func InitLogger() (*zap.Logger, error) {
	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
		OutputPaths:      []string{"stdout", "logs/app.log"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",
		},
	}

	return cfg.Build()
}

func LoadConfig() (config Config, err error) {

	viper.SetConfigFile("config/.env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	logger, err := InitLogger()
	if err != nil {
		return config, err
	}

	config.Logger = logger

	return config, err
}
