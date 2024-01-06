package application_commands

import (
	"context"
	"github.com/MichaelFraser99/discord-application-sdk/discord/config"
	"github.com/MichaelFraser99/discord-application-sdk/discord/model"
	"github.com/MichaelFraser99/discord-application-sdk/discord/utils"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/require"
	"net/http"
	"os"
	"testing"
)

var applicationId string
var commandId string
var service *ApplicationCommandService

func TestMain(m *testing.M) {
	id, ok := os.LookupEnv("APPLICATION_ID")
	if !ok {
		panic("no application ID present")
	} else {
		applicationId = id
	}
	token, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		panic("no bot token present")
	}

	service = New(config.NewConfig(config.TOKEN_TYPE_BOT, token, nil, http.DefaultClient))

	os.Exit(m.Run())
}

func TestApplicationCommandService_CreateCommand(t *testing.T) {
	tests := []struct {
		name     string
		request  *model.CreateApplicationCommand
		validate func(t *testing.T, command *model.ApplicationCommand, response *http.Response, err error)
	}{
		{
			name: "we can create a simple command",
			request: &model.CreateApplicationCommand{
				Name:         "simple-command",
				Description:  "a simple command",
				Type:         utils.Int(1),
				DmPermission: false,
				Nsfw:         false,
			},
			validate: func(t *testing.T, command *model.ApplicationCommand, response *http.Response, err error) {
				require.Nil(t, err)
				require.NotNil(t, response)
				require.Equal(t, 201, response.StatusCode)
				require.NotNil(t, command)
				assert.Equal(t, "simple-command", command.Name)
				assert.Equal(t, "a simple command", command.Description)
				assert.Equal(t, 1, command.Type)
				assert.Equal(t, false, command.DmPermission)
				assert.Equal(t, false, command.Nsfw)
				commandId = command.ID
			},
		},
		{
			name: "creating a duplicate command yields an error",
			request: &model.CreateApplicationCommand{
				Name:         "simple-command",
				Description:  "a simple command",
				Type:         utils.Int(1),
				DmPermission: false,
				Nsfw:         false,
			},
			validate: func(t *testing.T, command *model.ApplicationCommand, response *http.Response, err error) {
				require.NotNil(t, err)
				assert.Equal(t, nil, command)
				assert.Equal(t, nil, response)
				assert.Equal(t, "command with name simple-command already exists", err.Error())
			},
		},
		{
			name: "we can create a command with options",
			request: &model.CreateApplicationCommand{
				Name:        "command-options",
				Description: "a command with options",
				Type:        utils.Int(1),
				Options: []model.ApplicationCommandOption{
					{
						Name:        "add",
						Description: "add a new value",
						Type:        3,
						Required:    false,
					},
				},
				Nsfw:         true,
				DmPermission: true,
			},
			validate: func(t *testing.T, command *model.ApplicationCommand, response *http.Response, err error) {
				require.Nil(t, err)
				require.NotNil(t, response)
				require.Equal(t, 201, response.StatusCode)
				require.NotNil(t, command)
				assert.Equal(t, "command-options", command.Name)
				assert.Equal(t, "a command with options", command.Description)
				assert.Equal(t, 1, command.Type)
				assert.Equal(t, true, command.DmPermission)
				assert.Equal(t, true, command.Nsfw)
				require.Equal(t, 1, len(command.Options))
				assert.Equal(t, "add", command.Options[0].Name)
				assert.Equal(t, "add a new value", command.Options[0].Description)
				assert.Equal(t, 3, command.Options[0].Type)
				assert.Equal(t, false, command.Options[0].Required)
			},
		},
		{
			name: "we handle and display a bad request error from the API",
			request: &model.CreateApplicationCommand{
				Name:        "command-options",
				Description: "a command with options",
				Type:        utils.Int(3),
				Options: []model.ApplicationCommandOption{
					{
						Name:        "add",
						Description: "add a new value",
						Type:        3,
						Required:    false,
					},
				},
				Nsfw:         true,
				DmPermission: true,
			},
			validate: func(t *testing.T, command *model.ApplicationCommand, response *http.Response, err error) {
				require.NotNil(t, err)
				assert.Equal(t, nil, command)
				assert.Equal(t, nil, response)
				require.Contains(t, err.Error(), "error creating application command: Message: Invalid Form Body | Errors: ")
				require.Contains(t, err.Error(), "(options - Context menu commands cannot have options)")
				require.Contains(t, err.Error(), "(description - Context menu commands cannot have description)")
				assert.Equal(t, len("error creating application command: Message: Invalid Form Body | Errors: (options - Context menu commands cannot have options), (description - Context menu commands cannot have description)"), len(err.Error()))
			},
		},
		{
			name: "we handle and display a different bad request error from the API",
			request: &model.CreateApplicationCommand{
				Name:        "command with space And Capitals",
				Description: "a command with options",
				Type:        utils.Int(3),
				Options: []model.ApplicationCommandOption{
					{
						Name:        "Add Value",
						Description: "add a new value",
						Type:        3,
						Required:    false,
					},
				},
				Nsfw:         true,
				DmPermission: true,
			},
			validate: func(t *testing.T, command *model.ApplicationCommand, response *http.Response, err error) {
				require.NotNil(t, err)
				assert.Equal(t, nil, command)
				assert.Equal(t, nil, response)
				assert.Equal(t, "malformed request body", err.Error())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			command, response, err := service.CreateCommand(context.Background(), applicationId, tt.request)
			tt.validate(t, command, response, err)
		})
	}
}

func TestApplicationCommandService_PatchCommand(t *testing.T) {
	tests := []struct {
		name     string
		request  *model.PatchApplicationCommand
		validate func(t *testing.T, command *model.ApplicationCommand, response *http.Response, err error)
	}{
		{
			name: "we can update a commands description",
			request: &model.PatchApplicationCommand{
				Name:        "simple-command",
				Description: utils.String("this is an updated description"),
			},
			validate: func(t *testing.T, command *model.ApplicationCommand, response *http.Response, err error) {
				require.Nil(t, err)
				require.NotNil(t, response)
				require.Equal(t, 200, response.StatusCode)
				require.NotNil(t, command)
				assert.Equal(t, "simple-command", command.Name)
				assert.Equal(t, "this is an updated description", command.Description)
				assert.Equal(t, 1, command.Type)
				assert.Equal(t, false, command.DmPermission)
				assert.Equal(t, false, command.Nsfw)
			},
		},
		{
			name: "we cannot update a command to have a duplicate name",
			request: &model.PatchApplicationCommand{
				Name:        "command-options",
				Description: utils.String("this is an updated description"),
			},
			validate: func(t *testing.T, command *model.ApplicationCommand, response *http.Response, err error) {
				require.NotNil(t, err)
				require.Nil(t, response)
				require.Nil(t, command)
				assert.Equal(t, "error updating application command: Message: An application command with that name already exists", err.Error())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			command, response, err := service.PatchCommand(context.Background(), applicationId, commandId, tt.request)
			tt.validate(t, command, response, err)
		})
	}
}

func TestApplicationCommandService_DeleteCommand(t *testing.T) {
	commands, response, err := service.GetCommands(context.Background(), applicationId)
	require.Nil(t, err)
	require.Equal(t, 200, response.StatusCode)
	require.NotNil(t, commands)

	for _, c := range *commands {
		response, err = service.DeleteCommand(context.Background(), applicationId, c.ID)
		require.Nil(t, err)
		require.Equal(t, 204, response.StatusCode)
	}
}
