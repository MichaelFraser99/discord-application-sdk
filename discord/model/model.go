package model

import (
	"fmt"
	"strings"
)

type Application struct {
	Id          string  `json:"id"`
	Name        *string `json:"name"`
	Icon        any     `json:"icon"`
	Description *string `json:"description"`
	Type        any     `json:"type"`
	Bot         *struct {
		Id                   string  `json:"id"`
		Username             *string `json:"username"`
		Avatar               any     `json:"avatar"`
		Discriminator        *string `json:"discriminator"`
		PublicFlags          *int    `json:"public_flags"`
		PremiumType          *int    `json:"premium_type"`
		Flags                *int    `json:"flags"`
		Bot                  *bool   `json:"bot"`
		Banner               any     `json:"banner"`
		AccentColor          any     `json:"accent_color"`
		GlobalName           any     `json:"global_name"`
		AvatarDecorationData any     `json:"avatar_decoration_data"`
		BannerColor          any     `json:"banner_color"`
	} `json:"bot"`
	Summary                        *string `json:"summary"`
	BotPublic                      *bool   `json:"bot_public"`
	BotRequireCodeGrant            *bool   `json:"bot_require_code_grant"`
	VerifyKey                      *string `json:"verify_key"`
	Flags                          *int    `json:"flags"`
	Hook                           *bool   `json:"hook"`
	IsMonetized                    *bool   `json:"is_monetized"`
	RedirectUris                   *[]any  `json:"redirect_uris"`
	InteractionsEndpointUrl        *string `json:"interactions_endpoint_url"`
	RoleConnectionsVerificationUrl any     `json:"role_connections_verification_url"`
	Owner                          *struct {
		Id                   string  `json:"id"`
		Username             *string `json:"username"`
		Avatar               any     `json:"avatar"`
		Discriminator        *string `json:"discriminator"`
		PublicFlags          *int    `json:"public_flags"`
		PremiumType          *int    `json:"premium_type"`
		Flags                *int    `json:"flags"`
		Banner               any     `json:"banner"`
		AccentColor          any     `json:"accent_color"`
		GlobalName           *string `json:"global_name"`
		AvatarDecorationData any     `json:"avatar_decoration_data"`
		BannerColor          any     `json:"banner_color"`
	} `json:"owner"`
	ApproximateGuildCount        *int   `json:"approximate_guild_count"`
	InteractionsEventTypes       *[]any `json:"interactions_event_types"`
	InteractionsVersion          *int   `json:"interactions_version"`
	ExplicitContentFilter        *int   `json:"explicit_content_filter"`
	RpcApplicationState          *int   `json:"rpc_application_state"`
	StoreApplicationState        *int   `json:"store_application_state"`
	CreatorMonetizationState     *int   `json:"creator_monetization_state"`
	VerificationState            *int   `json:"verification_state"`
	IntegrationPublic            *bool  `json:"integration_public"`
	IntegrationRequireCodeGrant  *bool  `json:"integration_require_code_grant"`
	DiscoverabilityState         *int   `json:"discoverability_state"`
	DiscoveryEligibilityFlags    *int   `json:"discovery_eligibility_flags"`
	MonetizationState            *int   `json:"monetization_state"`
	MonetizationEligibilityFlags *int   `json:"monetization_eligibility_flags"`
	Team                         any    `json:"team"`
	InternalGuildRestriction     *int   `json:"internal_guild_restriction"`
}

type ApplicationCommand struct {
	ID                       string                     `json:"id"`
	Type                     int                        `json:"type,omitempty"`
	ApplicationID            string                     `json:"application_id,omitempty"`
	GuildID                  *string                    `json:"guild_id,omitempty"`
	Name                     string                     `json:"name"`
	NameLocalizations        *map[string]string         `json:"name_localizations,omitempty"`
	Description              string                     `json:"description"`
	DescriptionLocalizations *map[string]string         `json:"description_localizations,omitempty"`
	Options                  []ApplicationCommandOption `json:"options"` //When defining multiple, ensure required values are listed before optional values.
	DefaultMemberPermissions *string                    `json:"default_member_permissions,omitempty"`
	DmPermission             bool                       `json:"dm_permission"`
	DefaultPermission        *bool                      `json:"default_permission,omitempty"`
	Nsfw                     bool                       `json:"nsfw"`
	Version                  string                     `json:"version"`
}

type ApplicationCommandOption struct {
	Type                     int                              `json:"type" validate:"oneof=1 3 4 5 10"`
	Name                     string                           `json:"name" validate:"required,excludes= "`
	NameLocalizations        *map[string]string               `json:"name_localizations,omitempty"`
	Description              string                           `json:"description"`
	DescriptionLocalizations *map[string]string               `json:"description_localizations,omitempty"`
	Required                 bool                             `json:"required"`
	Choices                  []ApplicationCommandOptionChoice `json:"choices" validate:"dive"`
	Options                  []ApplicationCommandOption       `json:"options" validate:"dive"` //When defining multiple, ensure required values are listed before optional values.
	ChannelTypes             *[]int                           `json:"channel_types,omitempty"`
	MinValue                 *int                             `json:"min_value,omitempty"`
	MaxValue                 *int                             `json:"max_value,omitempty"`
	MinLength                *int                             `json:"min_length,omitempty"`
	MaxLength                *int                             `json:"max_length,omitempty"`
	AutoComplete             *bool                            `json:"autocomplete,omitempty"` //Must be false if choices are defined
}

type ApplicationCommandOptionChoice struct {
	Name              string             `json:"name" validate:"required,excludes= "`
	NameLocalizations *map[string]string `json:"name_localizations,omitempty"`
	Value             any                `json:"value"` //Can be string, integer, or double. Note: if string, max length is 100
}

type CreateApplicationCommand struct {
	Name                     string                     `json:"name" validate:"required,excludes= "`
	NameLocalizations        *map[string]string         `json:"name_localizations,omitempty"`
	Description              string                     `json:"description"`
	DescriptionLocalizations *map[string]string         `json:"description_localizations,omitempty"`
	Options                  []ApplicationCommandOption `json:"options" validate:"dive"` //When defining multiple, ensure required values are listed before optional values.
	DefaultMemberPermissions *string                    `json:"default_member_permissions,omitempty"`
	DmPermission             bool                       `json:"dm_permission"`
	DefaultPermission        *bool                      `json:"default_permission,omitempty"`
	Type                     *int                       `json:"type,omitempty"` //defaults to 1
	Nsfw                     bool                       `json:"nsfw"`
}

type PatchApplicationCommand struct {
	Name                     string                     `json:"name,omitempty" validate:"required,excludes= "`
	NameLocalizations        *map[string]string         `json:"name_localizations,omitempty"`
	Description              *string                    `json:"description,omitempty"`
	DescriptionLocalizations *map[string]string         `json:"description_localizations,omitempty"`
	Options                  []ApplicationCommandOption `json:"options" validate:"dive"` //When defining multiple, ensure required values are listed before optional values.
	DefaultMemberPermissions *string                    `json:"default_member_permissions,omitempty"`
	DmPermission             *bool                      `json:"dm_permission,omitempty"`
	DefaultPermission        *bool                      `json:"default_permission,omitempty"`
	Nsfw                     *bool                      `json:"nsfw,omitempty"`
}

type PatchApplication struct {
	InteractionsEndpointUrl *string `json:"interactions_endpoint_url"`
}

type ErrorResponse struct {
	Message    string            `json:"message"`
	RetryAfter *float32          `json:"retry_after"`
	Code       int               `json:"code"`
	Errors     map[string]Causes `json:"errors"`
}

type Causes struct {
	Errors []Error `json:"_errors"`
}
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e ErrorResponse) Format() string {
	if e.RetryAfter == nil {
		var causes []string
		for k, v := range e.Errors {
			var errorStrings []string
			for _, e := range v.Errors {
				errorStrings = append(errorStrings, e.Message)
			}
			causes = append(causes, fmt.Sprintf("(%s - %s)", k, strings.Join(errorStrings, ", ")))
		}
		if len(causes) > 0 {
			return fmt.Sprintf("Message: %s | Errors: %s", e.Message, strings.Join(causes, ", "))
		} else {
			return fmt.Sprintf("Message: %s", e.Message)
		}
	} else {
		return fmt.Sprintf("Message: %s | Retry after: %f", e.Message, *e.RetryAfter)
	}
}
