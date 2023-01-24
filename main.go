package main

import (
	"example/backend/db"
	"example/backend/logger"
	"example/backend/middleware"
	"example/backend/router"
	"example/backend/server"
	"net"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"

	"google.golang.org/grpc"
)

func runRestApi() {
	r := gin.New()

	r.Use(ginzap.Ginzap(logger.Logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger.Logger, true))

	r.Use(middleware.ErrorHandler)

	authMiddleware, err := jwt.New(middleware.AuthMiddleware)
	if err != nil {
		logger.Logger.Fatal(err.Error())
		panic(err)
	}
	r.Use(authMiddleware.MiddlewareFunc())

	router.AddRoutes(r)
	r.Run(":3001")
}

func runGrpc() {
	lis, err := net.Listen("tcp", ":3002")
	if err != nil {
		logger.Logger.Fatal(err.Error())
		panic(err)
	}

	s := grpc.NewServer()
	server.Register(s)
	logger.Logger.Sugar().Infof("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		logger.Logger.Fatal(err.Error())
		panic(err)
	}
}

func main() {
	db.Open()
	defer db.Close()

	go runRestApi()
	runGrpc()
}
