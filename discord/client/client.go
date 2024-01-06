package client

import (
	"github.com/MichaelFraser99/discord-application-sdk/discord/config"
	"github.com/MichaelFraser99/discord-application-sdk/services/application_commands"
	"github.com/MichaelFraser99/discord-application-sdk/services/applications"
)

type Client struct {
	ApplicationCommand application_commands.ApplicationCommandsAPI
	Application        applications.ApplicationAPI
}

func NewClient(cfg *config.Config) *Client {
	return &Client{
		ApplicationCommand: application_commands.New(cfg),
		Application:        applications.New(cfg),
	}
}
