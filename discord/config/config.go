package config

import (
	"fmt"
	"net/http"
)

type TokenType string

const (
	TOKEN_TYPE_BOT    TokenType = "Bot"
	TOKEN_TYPE_BEARER           = "Bearer"
)

type Config struct {
	TokenType  TokenType
	Token      string
	BaseUrl    string  `default:"https://discord.com/api"`
	ApiVersion *string // omitting will default to the current default version detailed within the discord API documentation
	HTTPClient *http.Client
}

func NewConfig(tokenType TokenType, token string, baseUrl string, apiVersion *string, httpClient *http.Client) *Config {
	return &Config{
		TokenType:  tokenType,
		Token:      token,
		BaseUrl:    baseUrl,
		ApiVersion: apiVersion,
		HTTPClient: httpClient,
	}
}

func (c *Config) GetVersionedUrl() string {
	if c.ApiVersion == nil {
		return c.BaseUrl
	} else {
		return fmt.Sprintf("%s/v%s", c.BaseUrl, *c.ApiVersion)
	}
}
