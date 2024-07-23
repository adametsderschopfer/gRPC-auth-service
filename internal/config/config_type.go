package config

import "time"

type Config struct {
	Env    string `yaml:"env" env-default:"local"`
	GRPC   `yaml:"grpc"`
	SQLite `yaml:"sqlite"`
}

type SQLite struct {
	StoragePath string `yaml:"storage_path" env-required:"true"`
}

type GRPC struct {
	Port    int           `yaml:"port" env-required:"true"`
	Timeout time.Duration `yaml:"timeout" env-required:"true"`
}
