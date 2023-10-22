package application_commands

import (
	"context"
	"encoding/json"
	"github.com/MichaelFraser99/discord-application-sdk/discord/config"
	"github.com/MichaelFraser99/discord-application-sdk/discord/model"
	"github.com/MichaelFraser99/discord-application-sdk/discord/utils"
	"net/http"
	"testing"
)

func TestTmp(t *testing.T) {
	service := New(config.NewConfig(config.TOKEN_TYPE_BOT, "MTEwNzM5MzU2MTE3MjkyMjQ5OQ.GiPREJ.9HcKwoTth3wVJMCS-4HPuBJxQRWWs5FJW4Yqb4", "https://discord.com/api", &http.Client{}))

	c, r, err := service.CreateCommand(context.Background(), "1107393561172922499", &model.CreateApplicationCommand{
		Name:        "test_command",
		Description: "a test command",
		Type:        utils.Int(1),
		Options: []model.ApplicationCommandOption{
			{
				Name:        "type3",
				Description: "type3 test",
				Type:        10,
				Required:    true,
				Choices: []model.ApplicationCommandOptionChoice{
					{
						Name:  "user1",
						Value: "1234567890",
					},
				},
			},
		},
	})
	//choices 3
	//no choices
	//bad

	if err != nil {
		t.Fatal(err)
	}

	if r.StatusCode != 201 {
		t.Fatal("expected status code 201")
	}

	cbytes, err := json.Marshal(c)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(cbytes))
}
