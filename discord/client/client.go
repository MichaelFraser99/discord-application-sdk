package client

import (
	"github.com/MichaelFraser99/discord-application-sdk/discord/config"
	"github.com/MichaelFraser99/discord-application-sdk/services/applicationCommands"
)

type Client struct {
	ApplicationCommand applicationCommands.ApplicationCommandsAPI
}

func (c *Client) NewClient(cfg *config.Config) *Client {
	return &Client{
		ApplicationCommand: applicationCommands.New(cfg),
	}
}
