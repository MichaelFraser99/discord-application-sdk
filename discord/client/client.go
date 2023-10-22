package client

import (
	"github.com/MichaelFraser99/discord-application-sdk/discord/config"
	"github.com/MichaelFraser99/discord-application-sdk/services/application_commands"
)

type Client struct {
	ApplicationCommand application_commands.ApplicationCommandsAPI
}

func NewClient(cfg *config.Config) *Client {
	return &Client{
		ApplicationCommand: application_commands.New(cfg),
	}
}
