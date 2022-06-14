package controller

import (
	"chatGo/src/infrastructure/keycloak"
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"io"
	"net/http"
	"time"
)

func auth(router *gin.Engine) {
	router.GET("/", authController)
	//router.GET("/auth/callback", callBack)
}

func authController(c *gin.Context) {
	state, err := randString(16)
	if err != nil {
		http.Error(c.Writer, "Internal error", http.StatusInternalServerError)
		return
	}
	nonce, err := randString(16)
	if err != nil {
		http.Error(c.Writer, "Internal error", http.StatusInternalServerError)
		return
	}
	setCallbackCookie(c.Writer, c.Request, "state", state, true)
	setCallbackCookie(c.Writer, c.Request, "nonce", nonce, true)
	c.Redirect(http.StatusFound, keycloak.ConfigKeyCloak.Oauth.AuthCodeURL(state, oidc.Nonce(nonce)))
}

func randString(n int) (string, error) {
	data := make([]byte, n)
	if _, err := io.ReadFull(rand.Reader, data); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

func setCallbackCookie(w http.ResponseWriter, r *http.Request, name, value string, httpOnly bool) {
	c := &http.Cookie{
		Name:     name,
		Value:    value,
		MaxAge:   int(time.Hour.Seconds()),
		Secure:   r.TLS != nil,
		HttpOnly: httpOnly,
	}
	http.SetCookie(w, c)
}

func callBack(c *gin.Context) {
	if c.Query("state") != "123" {
		http.Error(c.Writer, "State inválido", http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	token, err := keycloak.ConfigKeyCloak.Oauth.Exchange(ctx, c.Query("code"))
	if err != nil {
		http.Error(c.Writer, "Falha ao trocar o token", http.StatusInternalServerError)
	}

	idToken, ok := token.Extra("id_token").(string)
	if !ok {
		http.Error(c.Writer, "Falha ao gerar o IDToken", http.StatusInternalServerError)
		return
	}

	userInfo, err := keycloak.ConfigKeyCloak.Provider.UserInfo(ctx, oauth2.StaticTokenSource(token))
	if err != nil {
		http.Error(c.Writer, "Falha ao pegar UserInfo", http.StatusInternalServerError)
	}

	resp := struct {
		AccessToken *oauth2.Token
		IDToken     string
		UserInfo    *oidc.UserInfo
	}{
		token,
		idToken,
		userInfo,
	}

	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	c.Writer.Write(data)
}
