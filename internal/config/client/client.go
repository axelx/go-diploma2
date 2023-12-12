package client

import (
	"flag"
	"os"
)

type ConfigClient struct {
	RunAddr string
}

func NewConfigClient() *ConfigClient {
	conf := ConfigClient{
		RunAddr: "",
	}
	parseFlagsClient(&conf)
	return &conf
}

func parseFlagsClient(c *ConfigClient) {
	flag.StringVar(&c.RunAddr, "a", ":50051", "address and port to run server")

	flag.Parse()

	if envRunAddr := os.Getenv("ADDRESS"); envRunAddr != "" {
		c.RunAddr = envRunAddr
	}
}
