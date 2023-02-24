package rest

import (
	"errors"
	"example/backend/env"
	"fmt"

	gojwt "github.com/golang-jwt/jwt"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Jwt() echo.MiddlewareFunc {
	config := echojwt.Config{
		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
			keyFunc := func(token *gojwt.Token) (interface{}, error) {
				if token.Method.Alg() != env.CONFIG.JWT_ALG {
					return nil, fmt.Errorf("invalid signing algorithm: %v", token.Method.Alg())
				}
				return []byte(env.CONFIG.JWT_KEY), nil
			}

			token, err := gojwt.Parse(auth, keyFunc)
			if err != nil {
				return nil, err
			}
			if !token.Valid {
				return nil, errors.New("invalid token")
			}
			claims, _ := token.Claims.(gojwt.MapClaims)
			if !claims.VerifyIssuer(env.CONFIG.JWT_ISS, true) {
				return nil, fmt.Errorf("invalid iss claim: %v", claims["iss"])
			}
			if !claims.VerifyAudience(env.CONFIG.JWT_AUD, true) {
				return nil, fmt.Errorf("invalid aud claim: %v", claims["aud"])
			}

			return token, nil
		},
	}
	return echojwt.WithConfig(config)
}
