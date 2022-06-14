package keycloak

import (
	"chatGo/src/infrastructure/settings"
	"context"
	"fmt"
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

	providerHost := fmt.Sprintf("http://%s/auth/realms/chat", config.KeycloakHost)
	provider, err := oidc.NewProvider(ctx, providerHost)
	if err != nil {
		log.Fatal(err)
	}

	redirectURL := fmt.Sprintf("http://%s/auth/callback", config.MainHost)
	oauthConfig := oauth2.Config{
		ClientID:     config.KeycloakClientID,
		ClientSecret: config.KeycloakClientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  redirectURL,
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
