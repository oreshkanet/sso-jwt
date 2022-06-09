package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oreshkanet/sso-jwt/internal/service"
	"net/http"
)

type handler struct {
	saSvc service.SoftwareAuthService
	eaSvc service.ExternalClientAuthService
	uaSvc service.UserAuthService
}

func NewHandler(
	softwareAuth service.SoftwareAuthService,
	externalAuth service.ExternalClientAuthService,
	userAuth service.UserAuthService,
) *handler {
	return &handler{
		saSvc: softwareAuth,
		eaSvc: externalAuth,
		uaSvc: userAuth,
	}
}

func (h *handler) softwareAuth(c *gin.Context) {
	req := new(softwareAuthRequest)
	if err := c.BindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorResponse(
			fmt.Errorf("parse JSON: %v", err),
		))
		return
	}

	token, err := h.saSvc.Auth(c.Request.Context(),
		req.SoftwareId,
		req.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorResponse(
			fmt.Errorf("parse JSON: %v", err),
		))
	}

	c.AbortWithStatusJSON(http.StatusOK, &softwareAuthResponse{
		AccessToken: token,
	})
}

func (h *handler) externalClientAuth(c *gin.Context) {
	req := new(externalClientAuthRequest)
	if err := c.BindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorResponse(
			fmt.Errorf("parse JSON: %v", err),
		))
		return
	}

	token, err := h.eaSvc.Auth(c.Request.Context(),
		req.ClientId,
		req.ClientSecret,
		req.GrantType,
		req.Scope)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorResponse(
			fmt.Errorf("parse JSON: %v", err),
		))
	}

	c.AbortWithStatusJSON(http.StatusOK, &externalClientAuthResponse{
		AccessToken: token,
	})
}

func (h *handler) userAuth(c *gin.Context) {
	req := new(userAuthRequest)
	if err := c.BindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorResponse(
			fmt.Errorf("parse JSON: %v", err),
		))
		return
	}

	token, err := h.uaSvc.Auth(c.Request.Context(),
		req.UserId,
		req.UserPassword,
		req.Scope)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorResponse(
			fmt.Errorf("parse JSON: %v", err),
		))
	}

	c.AbortWithStatusJSON(http.StatusOK, &userAuthResponse{
		AccessToken: token,
	})
}
