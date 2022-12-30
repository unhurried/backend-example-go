package rest

import (
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
		// TODO Externalize configuration to .env file.
		const iss = "http://localhost:3002"
		const aud = "todo-api"

		claims := ginJwt.ExtractClaims(c)
		return claims["iss"] == iss && claims["aud"] == aud
	},
}
