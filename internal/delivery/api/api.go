package api

import (
	"github.com/oreshkanet/sso-jwt/pkg/tokenizer"
	"net/http"
)

type Api struct {
	srv  *http.Server
	auth *tokenizer.Tokenizer
	h    *handler
}
