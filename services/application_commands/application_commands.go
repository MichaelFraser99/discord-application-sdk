package application_commands

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

type ApplicationCommandService struct {
	Config *config.Config
	HTTP   *discordHttp.HTTP
}

func New(cfg *config.Config) *ApplicationCommandService {
	return &ApplicationCommandService{
		HTTP:   discordHttp.NewHTTP(*cfg),
		Config: cfg,
	}
}

func (s *ApplicationCommandService) GetCommands(ctx context.Context, applicationID string) (output *[]model.ApplicationCommand, resp *http.Response, err error) {

	httpRequest, err := http.NewRequest(discordHttp.METHOD_GET, fmt.Sprintf("%s/applications/%s/commands", s.Config.GetVersionedUrl(), applicationID), nil)
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

func (s *ApplicationCommandService) GetCommand(ctx context.Context, applicationID, commandID string) (output *model.ApplicationCommand, resp *http.Response, err error) {
	httpRequest, err := http.NewRequest(discordHttp.METHOD_GET, fmt.Sprintf("%s/applications/%s/commands/%s", s.Config.GetVersionedUrl(), applicationID, commandID), nil)
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

func (s *ApplicationCommandService) CreateCommand(ctx context.Context, applicationID string, request *model.CreateApplicationCommand) (output *model.ApplicationCommand, resp *http.Response, err error) {
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

	httpRequest, err := http.NewRequest(discordHttp.METHOD_POST, fmt.Sprintf("%s/applications/%s/commands", s.Config.GetVersionedUrl(), applicationID), bytes.NewReader(requestBytes))
	if err != nil {
		return nil, nil, err
	}

	response, err := s.HTTP.Do(ctx, httpRequest)
	if err != nil {
		return nil, nil, err
	}

	defer response.Body.Close()

	responseBytes, err := io.ReadAll(response.Body)

	if response.StatusCode != 201 {
		if response.StatusCode == 200 {
			return nil, nil, fmt.Errorf("command with name %s already exists", request.Name)
		}
		var errorResponse model.ErrorResponse
		err = json.Unmarshal(responseBytes, &errorResponse)
		if err != nil {
			return nil, nil, err
		}
		return nil, nil, fmt.Errorf("error creating application command: %s", errorResponse.Format())
	}
	err = json.Unmarshal(responseBytes, &output)
	if err != nil {
		return nil, response, err
	}

	return output, response, nil
}

func (s *ApplicationCommandService) PatchCommand(ctx context.Context, applicationID, commandID string, request *model.PatchApplicationCommand) (output *model.ApplicationCommand, resp *http.Response, err error) {
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

	httpRequest, err := http.NewRequest(discordHttp.METHOD_PATCH, fmt.Sprintf("%s/applications/%s/commands/%s", s.Config.GetVersionedUrl(), applicationID, commandID), bytes.NewReader(requestBytes))
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
		return nil, nil, fmt.Errorf("error updating application command: %s", errorResponse.Format())
	}

	err = json.Unmarshal(responseBytes, &output)
	if err != nil {
		return nil, response, err
	}

	return output, response, nil
}

func (s *ApplicationCommandService) DeleteCommand(ctx context.Context, applicationID, commandID string) (resp *http.Response, err error) {
	httpRequest, err := http.NewRequest(discordHttp.METHOD_DELETE, fmt.Sprintf("%s/applications/%s/commands/%s", s.Config.GetVersionedUrl(), applicationID, commandID), nil)
	if err != nil {
		return nil, err
	}

	response, err := s.HTTP.Do(ctx, httpRequest)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return response, nil
}
