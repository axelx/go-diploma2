package server

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

}
