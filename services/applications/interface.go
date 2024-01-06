package applications

import (
	"context"
	"github.com/MichaelFraser99/discord-application-sdk/discord/model"
	"net/http"
)

type ApplicationAPI interface {
	GetApplication(ctx context.Context, applicationID string) (output *model.Application, resp *http.Response, err error)
	PatchApplication(ctx context.Context, applicationID string, request *model.PatchApplicationCommand) (output *model.Application, resp *http.Response, err error)
}
