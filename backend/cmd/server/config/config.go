package config

import "testKwd/backend/libs/utils"

var (
	serverHostEnvName = "SERVER_HOST"
	serverPortEnvName = "SERVER_PORT"
)

type Config struct {
	HostServer string
	PortServer string
}

func ServerConfig() *Config {
	return &Config{
		HostServer: utils.TrimEnv(serverHostEnvName),
		PortServer: utils.TrimEnv(serverPortEnvName),
	}
}
