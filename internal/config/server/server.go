package server

import (
	"flag"
	"os"
)

type ConfigServer struct {
	RunAddr     string
	DatabaseDSN string
}

func NewConfigServer() *ConfigServer {
	conf := ConfigServer{
		RunAddr:     "",
		DatabaseDSN: "",
	}
	parseFlagsServer(&conf)
	return &conf
}

func parseFlagsServer(c *ConfigServer) {
	flag.StringVar(&c.RunAddr, "a", ":50051", "address and port to run server")
	flag.StringVar(&c.DatabaseDSN, "d", "postgres://user:password@localhost:5464/go-ya-gophkeeper", "DATABASE_DSN string")

	flag.Parse()

	if envRunAddr := os.Getenv("ADDRESS"); envRunAddr != "" {
		c.RunAddr = envRunAddr
	}

	if envDatabaseDSN := os.Getenv("DATABASE_DSN"); envDatabaseDSN != "" {
		c.DatabaseDSN = envDatabaseDSN
	}
}
