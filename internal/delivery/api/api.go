package api

import (
	"github.com/gin-gonic/gin"
	"github.com/oreshkanet/sso-jwt/internal/service"
	"github.com/oreshkanet/sso-jwt/pkg/tokenizer"
	"net/http"
)

type Api struct {
	srv   *http.Server
	token *tokenizer.Tokenizer
	h     *handler
}

type Config struct {
	Srv          *http.Server
	Token        *tokenizer.Tokenizer
	SoftwareAuth service.SoftwareAuthService
	ExternalAuth service.ExternalClientAuthService
	UserAuth     service.UserAuthService
}

func NewApi(config *Config) *Api {
	return &Api{
		srv:   config.Srv,
		token: config.Token,
		h:     NewHandler(config.SoftwareAuth, config.ExternalAuth, config.UserAuth),
	}
}

func (a *Api) Run() error {
	// Создаём новый роутер
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	a.srv.Handler = router

	// Роутим эндпойнт авторизации
	ar := router.Group("/software")
	ar.POST("/auth", a.h.softwareAuth)

	// Роутим эндпойнт управление профилем пользователя
	ur := router.Group("/user")
	//userRouter.Use(a.h.UserMiddleware)
	//userRouter.POST("/change_role", a.h.userChangeRole)
	ur.PUT("/user", a.h.userAuth)

	// Роутим эндпойнт управление профилем пользователя
	ecr := router.Group("/external")
	ecr.PUT("/auth", a.h.externalClientAuth)

	return a.srv.ListenAndServe()
}

func (a *Api) Stop() error {
	return a.srv.Close()
}
