package applications

import (
	"context"
	"github.com/MichaelFraser99/discord-application-sdk/discord/config"
	"github.com/MichaelFraser99/discord-application-sdk/discord/model"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/require"
	"net/http"
	"os"
	"testing"
)

var applicationId string
var service *ApplicationService

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

func TestApplicationService_GetApplication(t *testing.T) {
	tests := []struct {
		name     string
		id       string
		validate func(t *testing.T, application *model.Application, response *http.Response, err error)
	}{
		{
			name: "we can get an application",
			id:   applicationId,
			validate: func(t *testing.T, application *model.Application, response *http.Response, err error) {
				require.Nil(t, err)
				require.NotNil(t, response)
				require.Equal(t, 200, response.StatusCode)
				require.NotNil(t, application)

				assert.Equal(t, "1193216970133356637", application.Id)
				assert.Equal(t, "SDK Unit Tests", application.Name)
				assert.Equal(t, nil, application.Icon)
				assert.Equal(t, "", application.Description)
				assert.Equal(t, nil, application.Type)
				assert.Equal(t, "1193216970133356637", application.Bot.Id)
				assert.Equal(t, "SDK Unit Tests", application.Bot.Username)
				assert.Equal(t, nil, application.Bot.Avatar)
				assert.Equal(t, "0468", application.Bot.Discriminator)
				assert.Equal(t, 0, application.Bot.PublicFlags)
				assert.Equal(t, 0, application.Bot.PremiumType)
				assert.Equal(t, 0, application.Bot.Flags)
				assert.Equal(t, true, application.Bot.Bot)
				assert.Equal(t, nil, application.Bot.Banner)
				assert.Equal(t, nil, application.Bot.AccentColor)
				assert.Equal(t, nil, application.Bot.GlobalName)
				assert.Equal(t, nil, application.Bot.AvatarDecorationData)
				assert.Equal(t, nil, application.Bot.BannerColor)
				assert.Equal(t, "", application.Summary)
				assert.Equal(t, true, application.BotPublic)
				assert.Equal(t, false, application.BotRequireCodeGrant)
				assert.Equal(t, "6927bbaa829db9dda34228e82e5350528f0384c6e26a4f319ad3ba62040350b0", application.VerifyKey)
				assert.Equal(t, 0, application.Flags)
				assert.Equal(t, true, application.Hook)
				assert.Equal(t, false, application.IsMonetized)
				assert.Equal(t, []interface{}{}, application.RedirectUris)
				assert.Equal(t, nil, application.InteractionsEndpointUrl)
				assert.Equal(t, nil, application.RoleConnectionsVerificationUrl)
				assert.Equal(t, "264095754413408257", application.Owner.Id)
				assert.Equal(t, "mmyyhhtt", application.Owner.Username)
				assert.Equal(t, nil, application.Owner.Avatar)
				assert.Equal(t, "0", application.Owner.Discriminator)
				assert.Equal(t, 0, application.Owner.PublicFlags)
				assert.Equal(t, 0, application.Owner.PremiumType)
				assert.Equal(t, 0, application.Owner.Flags)
				assert.Equal(t, nil, application.Owner.Banner)
				assert.Equal(t, nil, application.Owner.AccentColor)
				assert.Equal(t, "Mmyyhhtt", application.Owner.GlobalName)
				assert.Equal(t, nil, application.Owner.AvatarDecorationData)
				assert.Equal(t, nil, application.Owner.BannerColor)
				assert.Equal(t, 0, application.ApproximateGuildCount)
				assert.Equal(t, []interface{}{}, application.InteractionsEventTypes)
				assert.Equal(t, 1, application.InteractionsVersion)
				assert.Equal(t, 0, application.ExplicitContentFilter)
				assert.Equal(t, 0, application.RpcApplicationState)
				assert.Equal(t, 1, application.StoreApplicationState)
				assert.Equal(t, 1, application.CreatorMonetizationState)
				assert.Equal(t, 1, application.VerificationState)
				assert.Equal(t, true, application.IntegrationPublic)
				assert.Equal(t, false, application.IntegrationRequireCodeGrant)
				assert.Equal(t, nil, application.Team)
				assert.Equal(t, 2, application.InternalGuildRestriction)
			},
		},
		{
			name: "we don't crash on not found application",
			id:   "123456789",
			validate: func(t *testing.T, application *model.Application, response *http.Response, err error) {
				require.Equal(t, 404, response.StatusCode)
				require.Nil(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			application, response, err := service.GetApplication(context.Background(), tt.id)
			tt.validate(t, application, response, err)
		})
	}
}
