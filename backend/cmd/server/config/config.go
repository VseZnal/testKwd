package config

import "testKwd/libs/utils"

var (
	serverHostEnvName   = "SERVER_HOST"
	serverPortEnvName   = "SERVER_PORT"
	pgConnStringEnvName = "PG_CONNECTION_STRING"
)

type Config struct {
	HostServer   string
	PortServer   string
	PgConnString string
}

func ServerConfig() *Config {
	return &Config{
		HostServer:   utils.TrimEnv(serverHostEnvName),
		PortServer:   utils.TrimEnv(serverPortEnvName),
		PgConnString: utils.TrimEnv(pgConnStringEnvName),
	}
}
