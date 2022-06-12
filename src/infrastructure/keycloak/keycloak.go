package keycloak

import (
	"chatGo/src/infrastructure/settings"
	"context"
	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
	"log"
)

type configKeycloak struct {
	Oauth      oauth2.Config
	Provider   *oidc.Provider
	OidcConfig *oidc.Config
}

var ConfigKeyCloak configKeycloak

func Start(config *settings.GlobalConfig) {
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, "http://localhost:8080/auth/realms/chat")
	if err != nil {
		log.Fatal(err)
	}

	oauthConfig := oauth2.Config{
		ClientID:     config.KeycloakClientID,
		ClientSecret: config.KeycloakClientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  "http://localhost:8081/auth/callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email", "roles"},
	}

	oidcConfig := &oidc.Config{
		ClientID: config.KeycloakClientID,
	}

	ConfigKeyCloak = configKeycloak{
		Oauth:      oauthConfig,
		Provider:   provider,
		OidcConfig: oidcConfig,
	}
}
