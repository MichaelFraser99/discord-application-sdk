package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/MichaelFraser99/discord-application-sdk/discord/config"
	"net/http"
)

type HTTPMethods string

const (
	METHOD_GET    = "GET"
	METHOD_POST   = "POST"
	METHOD_PUT    = "PUT"
	METHOD_PATCH  = "PATCH"
	METHOD_DELETE = "DELETE"
)

type HTTP struct {
	Config config.Config
}

func NewHTTP(config config.Config) *HTTP {
	return &HTTP{
		Config: config,
	}
}

func (h *HTTP) Do(ctx context.Context, request *http.Request) (*http.Response, error) {
	if request == nil {
		return nil, errors.New("cannot perform request without a request object")
	}
	if h.Config.HTTPClient == nil {
		return nil, errors.New("cannot perform request without a http client")
	}

	request = request.WithContext(ctx)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("%s %s", h.Config.TokenType, h.Config.Token))

	response, err := h.Config.HTTPClient.Do(request)
	return response, err
}
