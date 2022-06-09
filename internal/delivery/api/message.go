package api

import (
	"fmt"
)

type ErrorResponse struct {
	ErrorMsg string `json:"err_msg"`
}

func newErrorResponse(err error) *ErrorResponse {
	return &ErrorResponse{
		fmt.Sprintf("%s", err.Error()),
	}
}

type softwareAuthRequest struct {
	SoftwareId string `json:"software_id"`
	Password   string `json:"password"`
}

type softwareAuthResponse struct {
	AccessToken string `json:"access_token"`
}

type externalClientAuthRequest struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	Scope        string `json:"scope"`
}

type externalClientAuthResponse struct {
	AccessToken string `json:"access_token"`
}

type userAuthRequest struct {
	UserId       string `json:"user_id"`
	UserPassword string `json:"password"`
	Scope        string `json:"scope"`
}

type userAuthResponse struct {
	AccessToken string `json:"access_token"`
}
