package api

import (
	"fmt"
)

type ListenConfig struct {
	Interface string
	Port      int
}

func (this *ListenConfig) Address() string {
	return fmt.Sprintf("%v:%v", this.Interface, this.Port)
}

type Config struct {
	Grpc        ListenConfig
	Http        ListenConfig
	SwaggerFile string
}

func NewConfig(listenInterface string, grpcPort int, httpPort int, swagger string) *Config {
	return &Config{
		Grpc:        ListenConfig{listenInterface, grpcPort},
		Http:        ListenConfig{listenInterface, httpPort},
		SwaggerFile: swagger,
	}
}
