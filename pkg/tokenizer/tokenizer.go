package tokenizer

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Tokenizer struct {
	signingKey     string
	expireDuration time.Duration
	signingMethod  *jwt.SigningMethodHMAC
}

func NewTokenizer(signingKey string, expireDuration time.Duration) *Tokenizer {
	return &Tokenizer{
		signingKey:     signingKey,
		expireDuration: expireDuration,
		signingMethod:  jwt.SigningMethodHS256,
	}
}

func (t *Tokenizer) Generate(issuer string, claims Claims) (string, error) {
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(t.expireDuration))
	claims.IssuedAt = jwt.NewNumericDate(time.Now())
	claims.Issuer = issuer

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)
	return token.SignedString(t.signingKey)
}

func (t *Tokenizer) ParseToken(accessToken string, claims *Claims) error {
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return t.signingKey, nil
	})

	if err != nil {
		return err
	}

	var ok bool
	if claims, ok = token.Claims.(*Claims); ok && token.Valid {
		return nil
	}

	return fmt.Errorf("invalid access token")
}
