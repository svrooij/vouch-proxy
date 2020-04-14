package homeassistant

import (
	"net/http"

	"github.com/vouch/vouch-proxy/handlers/common"
	"github.com/vouch/vouch-proxy/pkg/cfg"
	"github.com/vouch/vouch-proxy/pkg/structs"
	"go.uber.org/zap"
)

// Provider provider specific functions
type Provider struct{}

var log *zap.SugaredLogger

// Configure see main.go configure()
func (Provider) Configure() {
	log = cfg.Cfg.Logger
}

// GetUserInfo provider specific call to get userinfomation
// More info: https://developers.home-assistant.io/docs/en/auth_api.html
func (Provider) GetUserInfo(r *http.Request, user *structs.User, customClaims *structs.CustomClaims, ptokens *structs.PTokens) (rerr error) {
	_, providerToken, err := common.PrepareTokensAndClient(r, ptokens, false)
	if err != nil {
		return err
	}
	ptokens.PAccessToken = providerToken.Extra("access_token").(string)
	// Home assistant does not provide an API to query username, so we statically set it to "homeassistant"
	user.Username = "homeassistant"
	return nil
}