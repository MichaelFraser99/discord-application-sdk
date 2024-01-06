package applications

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/MichaelFraser99/discord-application-sdk/discord/config"
	discordHttp "github.com/MichaelFraser99/discord-application-sdk/discord/http"
	"github.com/MichaelFraser99/discord-application-sdk/discord/model"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
)

type ApplicationService struct {
	Config *config.Config
	HTTP   *discordHttp.HTTP
}

func New(cfg *config.Config) *ApplicationService {
	return &ApplicationService{
		HTTP:   discordHttp.NewHTTP(*cfg),
		Config: cfg,
	}
}

func (s *ApplicationService) GetApplication(ctx context.Context, applicationID string) (output *model.Application, resp *http.Response, err error) {
	httpRequest, err := http.NewRequest(discordHttp.METHOD_GET, fmt.Sprintf("%s/applications/%s", s.Config.GetVersionedUrl(), applicationID), nil)
	if err != nil {
		return nil, nil, err
	}

	response, err := s.HTTP.Do(ctx, httpRequest)
	if err != nil {
		return nil, nil, err
	}

	defer response.Body.Close()

	responseBytes, err := io.ReadAll(response.Body)

	err = json.Unmarshal(responseBytes, &output)
	if err != nil {
		return nil, response, err
	}

	return output, response, nil
}

func (s *ApplicationService) PatchApplication(ctx context.Context, applicationID string, request *model.PatchApplication) (output *model.Application, resp *http.Response, err error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request cannot be nil")
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(*request)
	if err != nil {
		return nil, nil, fmt.Errorf("malformed request body")
	}

	requestBytes, err := json.Marshal(request)
	if err != nil {
		return nil, nil, err
	}

	httpRequest, err := http.NewRequest(discordHttp.METHOD_PATCH, fmt.Sprintf("%s/applications/%s", s.Config.GetVersionedUrl(), applicationID), bytes.NewReader(requestBytes))
	if err != nil {
		return nil, nil, err
	}

	response, err := s.HTTP.Do(ctx, httpRequest)
	if err != nil {
		return nil, nil, err
	}

	defer response.Body.Close()

	responseBytes, err := io.ReadAll(response.Body)

	if response.StatusCode != 200 {
		var errorResponse model.ErrorResponse
		err = json.Unmarshal(responseBytes, &errorResponse)
		if err != nil {
			return nil, nil, err
		}
		return nil, nil, fmt.Errorf("error updating application: %s", errorResponse.Format())
	}

	err = json.Unmarshal(responseBytes, &output)
	if err != nil {
		return nil, response, err
	}

	return output, response, nil
}
