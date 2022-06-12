package middleware

import (
	"chatGo/src/infrastructure/keycloak"
	"context"
	"encoding/json"
	"github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"net/http"
)

func ValidateToken(c *gin.Context) {
	err, done := validateState(c)
	if done {
		http.Error(c.Writer, "State not valid", http.StatusBadRequest)
		c.AbortWithError(404, err)
	}

	ctx := context.Background()
	oauth2Token, err := keycloak.ConfigKeyCloak.Oauth.Exchange(ctx, c.Query("code"))
	if err != nil {
		http.Error(c.Writer, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		c.AbortWithError(404, err)
	}
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		http.Error(c.Writer, "No id_token field in oauth2 token.", http.StatusInternalServerError)
		c.AbortWithError(404, err)
	}

	verifier := keycloak.ConfigKeyCloak.Provider.Verifier(keycloak.ConfigKeyCloak.OidcConfig)
	idToken, err := verifier.Verify(ctx, rawIDToken)
	if err != nil {
		http.Error(c.Writer, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
		c.AbortWithError(404, err)
	}

	if nonceValidation(c, err, idToken) {
		http.Error(c.Writer, "Nonce not valid", http.StatusBadRequest)
		c.AbortWithError(404, err)
	}

	oauth2Token.AccessToken = "*REDACTED*"

	resp := struct {
		OAuth2Token   *oauth2.Token
		IDTokenClaims *json.RawMessage // ID Token payload is just JSON.
	}{oauth2Token, new(json.RawMessage)}

	if err := idToken.Claims(&resp.IDTokenClaims); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		c.AbortWithError(404, err)
	}

	// Extract custom claims
	var claims struct {
		PreferredUsername string `json:"preferred_username"`
	}
	if err := idToken.Claims(&claims); err != nil {
		// handle error
	}
	c.Set("preferred_username", claims.PreferredUsername)

}

func nonceValidation(c *gin.Context, err error, idToken *oidc.IDToken) bool {
	nonce, err := c.Request.Cookie("nonce")
	if err != nil {
		return true
	}
	if idToken.Nonce != nonce.Value {
		return true
	}
	return false
}

func validateState(c *gin.Context) (error, bool) {
	state, err := c.Request.Cookie("state")
	if err != nil {
		return nil, true
	}
	if c.Query("state") != state.Value {
		return nil, true
	}
	return err, false
}
