package client

type ConfigClient struct {
	RunAddr string
}

func NewConfigClient() *ConfigClient {
	conf := ConfigClient{
		RunAddr: "",
	}
	parseFlagsServer(&conf)
	return &conf
}

func parseFlagsServer(c *ConfigClient) {

}
