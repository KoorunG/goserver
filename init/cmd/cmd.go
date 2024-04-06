package cmd

import (
	"goserver/config"
	"goserver/network"
)

type Cmd struct {
	config  *config.Config
	network *network.Network
}

func NewCmd(filepath string) *Cmd {
	c := &Cmd{
		config:  config.NewConfig(filepath),
		network: network.NewNetwork(),
	}

	c.network.ServerStart(c.config.Server.Port)
	return c
}
