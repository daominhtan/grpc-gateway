package config

type ServerConfig struct {
	Network string
	Address string
}

var DefaultGRPCServerConfig = ServerConfig{
	Network: "tcp",
	Address: "localhost:9090",
}

var DefaultReverseConfig = ServerConfig{
	Address: "localhost:8080",
}
