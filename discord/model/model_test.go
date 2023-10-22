package model

import (
	"github.com/MichaelFraser99/discord-application-sdk/discord/utils"
	"github.com/go-playground/validator/v10"
	"testing"
)

func Test_CreateApplicationCommand(t *testing.T) {
	tests := []struct {
		name   string
		input  CreateApplicationCommand
		result *string
	}{
		{
			name: "valid",
			input: CreateApplicationCommand{
				Name:              "test",
				Description:       "test",
				DmPermission:      false,
				DefaultPermission: utils.Bool(false),
				Type:              utils.Int(3),
				Nsfw:              false,
			},
			result: nil,
		},
		{
			name: "invalid - no space in name",
			input: CreateApplicationCommand{
				Name:              "test test",
				Description:       "test",
				DmPermission:      false,
				DefaultPermission: utils.Bool(false),
				Type:              utils.Int(3),
				Nsfw:              false,
			},
			result: utils.String("Key: 'CreateApplicationCommand.Name' Error:Field validation for 'Name' failed on the 'excludes' tag"),
		},
		{
			name: "invalid - no space in option name",
			input: CreateApplicationCommand{
				Name:              "test",
				Description:       "test",
				DmPermission:      false,
				DefaultPermission: utils.Bool(false),
				Options: []ApplicationCommandOption{
					{
						Name: "test test",
						Type: 3,
					},
				},
				Type: utils.Int(3),
				Nsfw: false,
			},
			result: utils.String("Key: 'CreateApplicationCommand.Options[0].Name' Error:Field validation for 'Name' failed on the 'excludes' tag"),
		},
		{
			name: "invalid - name required",
			input: CreateApplicationCommand{
				Description:       "test",
				DmPermission:      false,
				DefaultPermission: utils.Bool(false),
				Type:              utils.Int(3),
				Nsfw:              false,
			},
			result: utils.String("Key: 'CreateApplicationCommand.Name' Error:Field validation for 'Name' failed on the 'required' tag"),
		},
		{
			name: "valid - option type 1",
			input: CreateApplicationCommand{
				Name:              "test",
				Description:       "test",
				DmPermission:      false,
				DefaultPermission: utils.Bool(false),
				Type:              utils.Int(3),
				Nsfw:              false,
				Options: []ApplicationCommandOption{
					{
						Type: 1,
						Name: "test",
					},
				},
			},
			result: nil,
		},
		{
			name: "invalid - option type 2",
			input: CreateApplicationCommand{
				Name:              "test",
				Description:       "test",
				DmPermission:      false,
				DefaultPermission: utils.Bool(false),
				Type:              utils.Int(3),
				Nsfw:              false,
				Options: []ApplicationCommandOption{
					{
						Type: 2,
						Name: "test",
					},
				},
			},
			result: utils.String("Key: 'CreateApplicationCommand.Options[0].Type' Error:Field validation for 'Type' failed on the 'oneof' tag"),
		},
		{
			name: "valid - option type 3",
			input: CreateApplicationCommand{
				Name:              "test",
				Description:       "test",
				DmPermission:      false,
				DefaultPermission: utils.Bool(false),
				Type:              utils.Int(3),
				Nsfw:              false,
				Options: []ApplicationCommandOption{
					{
						Type: 3,
						Name: "test",
					},
				},
			},
			result: nil,
		},
		{
			name: "valid - option type 4",
			input: CreateApplicationCommand{
				Name:              "test",
				Description:       "test",
				DmPermission:      false,
				DefaultPermission: utils.Bool(false),
				Type:              utils.Int(3),
				Nsfw:              false,
				Options: []ApplicationCommandOption{
					{
						Type: 4,
						Name: "test",
					},
				},
			},
			result: nil,
		},
		{
			name: "valid - option type 5",
			input: CreateApplicationCommand{
				Name:              "test",
				Description:       "test",
				DmPermission:      false,
				DefaultPermission: utils.Bool(false),
				Type:              utils.Int(3),
				Nsfw:              false,
				Options: []ApplicationCommandOption{
					{
						Type: 5,
						Name: "test",
					},
				},
			},
			result: nil,
		},
		{
			name: "invalid - option type 6",
			input: CreateApplicationCommand{
				Name:              "test",
				Description:       "test",
				DmPermission:      false,
				DefaultPermission: utils.Bool(false),
				Type:              utils.Int(3),
				Nsfw:              false,
				Options: []ApplicationCommandOption{
					{
						Type: 6,
						Name: "test",
					},
				},
			},
			result: utils.String("Key: 'CreateApplicationCommand.Options[0].Type' Error:Field validation for 'Type' failed on the 'oneof' tag"),
		},
		{
			name: "invalid - option type 7",
			input: CreateApplicationCommand{
				Name:              "test",
				Description:       "test",
				DmPermission:      false,
				DefaultPermission: utils.Bool(false),
				Type:              utils.Int(3),
				Nsfw:              false,
				Options: []ApplicationCommandOption{
					{
						Type: 7,
						Name: "test",
					},
				},
			},
			result: utils.String("Key: 'CreateApplicationCommand.Options[0].Type' Error:Field validation for 'Type' failed on the 'oneof' tag"),
		},
		{
			name: "invalid - option type 8",
			input: CreateApplicationCommand{
				Name:              "test",
				Description:       "test",
				DmPermission:      false,
				DefaultPermission: utils.Bool(false),
				Type:              utils.Int(3),
				Nsfw:              false,
				Options: []ApplicationCommandOption{
					{
						Type: 8,
						Name: "test",
					},
				},
			},
			result: utils.String("Key: 'CreateApplicationCommand.Options[0].Type' Error:Field validation for 'Type' failed on the 'oneof' tag"),
		},
		{
			name: "invalid - option type 9",
			input: CreateApplicationCommand{
				Name:              "test",
				Description:       "test",
				DmPermission:      false,
				DefaultPermission: utils.Bool(false),
				Type:              utils.Int(3),
				Nsfw:              false,
				Options: []ApplicationCommandOption{
					{
						Type: 9,
						Name: "test",
					},
				},
			},
			result: utils.String("Key: 'CreateApplicationCommand.Options[0].Type' Error:Field validation for 'Type' failed on the 'oneof' tag"),
		},
		{
			name: "valid - option type 10",
			input: CreateApplicationCommand{
				Name:              "test",
				Description:       "test",
				DmPermission:      false,
				DefaultPermission: utils.Bool(false),
				Type:              utils.Int(3),
				Nsfw:              false,
				Options: []ApplicationCommandOption{
					{
						Type: 10,
						Name: "test",
					},
				},
			},
			result: nil,
		},
		{
			name: "invalid - option type 11",
			input: CreateApplicationCommand{
				Name:              "test",
				Description:       "test",
				DmPermission:      false,
				DefaultPermission: utils.Bool(false),
				Type:              utils.Int(3),
				Nsfw:              false,
				Options: []ApplicationCommandOption{
					{
						Type: 11,
						Name: "test",
					},
				},
			},
			result: utils.String("Key: 'CreateApplicationCommand.Options[0].Type' Error:Field validation for 'Type' failed on the 'oneof' tag"),
		},
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := validate.Struct(test.input)
			if test.result != nil || err != nil {
				if test.result == nil && err != nil {
					t.Errorf("Expected no error, got %v", err)
				} else if test.result != nil && err == nil {
					t.Errorf("Expected %v, got no error", *test.result)
				}

				if err.Error() != *test.result {
					t.Errorf("Expected %v, got %v", *test.result, err)
				}
			}
		})
	}
}
