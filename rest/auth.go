package rest

import (
	"example/backend/env"

	ginJwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

var AuthMiddleware = &ginJwt.GinJWTMiddleware{
	SigningAlgorithm: "HS256",
	Key:              []byte("secret"),
	ParseOptions: []jwt.ParserOption{func(p *jwt.Parser) {
	}},
	Authorizator: func(_ interface{}, c *gin.Context) bool {
		iss := env.CONFIG.JWT_ISS
		aud := env.CONFIG.JWT_AUD

		claims := ginJwt.ExtractClaims(c)
		return claims["iss"] == iss && claims["aud"] == aud
	},
}
