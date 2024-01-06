package config

import (
	"fmt"
	"github.com/MichaelFraser99/discord-application-sdk/discord/utils"
	"net/http"
)

type TokenType string

const (
	TOKEN_TYPE_BOT    TokenType = "Bot"
	TOKEN_TYPE_BEARER TokenType = "Bearer"
)

type Config struct {
	TokenType  TokenType
	Token      string
	BaseUrl    string
	apiVersion *string // omitting will default to the current default version detailed within the discord API documentation
	HTTPClient *http.Client
}

func NewConfig(tokenType TokenType, token string, baseUrl *string, httpClient *http.Client) *Config {
	var url string
	if baseUrl == nil {
		url = "https://discord.com/api"
	} else {
		url = *baseUrl
	}
	return &Config{
		TokenType:  tokenType,
		Token:      token,
		BaseUrl:    url,
		apiVersion: utils.String("10"),
		HTTPClient: httpClient,
	}
}

func (c *Config) GetVersionedUrl() string {
	return fmt.Sprintf("%s/v%s", c.BaseUrl, *c.apiVersion)
}
