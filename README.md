# Discord Application Commands SDK
An SDK enabling the programmatic use of the discord application commands API

## Usage
```go
package main

import (
	"net/http"

	"github.com/MichaelFraser99/discord-application-sdk/discord/config"
	"github.com/MichaelFraser99/discord-application-sdk/discord/client"
)

func main() {
	sdkConfig := config.NewConfig(
		config.TOKEN_TYPE_BOT,
		"*discord api token*",
		"https://discord.com/api",
		http.DefaultClient,
	)

	sdkClient := client.NewClient(sdkConfig)
}
```
