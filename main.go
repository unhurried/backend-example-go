package main

import (
	"example/backend/db"
	"example/backend/logger"
	"example/backend/rest"
	"example/backend/server"
	"net"

	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
)

func runRestApi() {
	s := rest.Server{}
	e := echo.New()
	e.Use(echomw.Logger())
	e.Use(echomw.Recover())
	e.Use(rest.Jwt())

	rest.RegisterHandlers(e, s)
	e.Logger.Fatal(e.Start(":3001"))
}

func runGrpc() {
	lis, err := net.Listen("tcp", ":3002")
	if err != nil {
		logger.Logger.Fatal(err.Error())
		panic(err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(server.Auth)))
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
