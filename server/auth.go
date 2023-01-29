package server

import (
	"context"
	"example/backend/env"

	gjwt "github.com/golang-jwt/jwt/v4"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
)

func Auth(ctx context.Context) (context.Context, error) {
	tokenString, err := grpc_auth.AuthFromMD(ctx, "Bearer")
	if err != nil {
		return ctx, err
	}

	token, err := gjwt.Parse(tokenString, func(token *gjwt.Token) (interface{}, error) {
		if token.Method.Alg() != env.CONFIG.JWT_ALG {
			return nil, gjwt.ErrTokenSignatureInvalid
		}
		return []byte(env.CONFIG.JWT_KEY), nil
	})
	if err != nil {
		return nil, err
	}

	claims := token.Claims.(gjwt.MapClaims)
	if !claims.VerifyIssuer(env.CONFIG.JWT_ISS, true) {
		return nil, gjwt.ErrTokenInvalidIssuer
	}
	if !claims.VerifyAudience(env.CONFIG.JWT_AUD, true) {
		return nil, gjwt.ErrTokenInvalidAudience
	}

	return ctx, nil
}
