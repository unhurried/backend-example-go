# backend-example-go

A simple Go project that serves REST APIs for TODO web app.

## Covered Features

* REST API routing with [echo](https://echo.labstack.com/).
* Code generation with [OpenAPI](https://www.openapis.org/) document and [oapi-codegen](https://github.com/deepmap/oapi-codegen).
* Request validation with [oapi-codegen/middleware](https://pkg.go.dev/github.com/deepmap/oapi-codegen/pkg/middleware).
* gRPC routing and code generation with [protobuf](https://pkg.go.dev/google.golang.org/protobuf)
* API authentication by JWT access tokens with [echo JWT middleware](https://echo.labstack.com/middleware/jwt/).
* Logging with [uber-go/zap](https://pkg.go.dev/go.uber.org/zap)
* Load environment variables and .env file with [godotenv](https://pkg.go.dev/github.com/joho/godotenv) and [env/v6](https://pkg.go.dev/github.com/caarlos0/env/v6)

\* For [v0.7.0](https://github.com/unhurried/backend-example-go/tree/v0.7.0) and earlier:

* API routing, (de)serialization and validation of request/response with [Gin](https://gin-gonic.com/).
* Code generation with [OpenAPI](https://www.openapis.org/) document and [OpenAPI Generator](https://openapi-generator.tech/).
* API authentication by JWT access tokens with [gin-jwt/v2](https://pkg.go.dev/github.com/appleboy/gin-jwt/v2).
* Logging with [uber-go/zap](https://pkg.go.dev/go.uber.org/zap)
* Load environment variables and .env file with [godotenv](https://pkg.go.dev/github.com/joho/godotenv) and [env/v6](https://pkg.go.dev/github.com/caarlos0/env/v6)

## Development Guide

You can run and build the project with `go` command.

```bash
# Run - http server will start and listen on port 3001
go run main.go
# Build - backend(.exe) will be created at the project root directory
go build
```

Also, [`launch.json`](./.vscode/launch.json) for Visual Studio Code is included in the project directory.

### OpenAPI specification

The specification of APIs are written in `openapi.yaml` using OpenAPI 3.0.3.

Model and router code can be generated with [oapi-generator](https://github.com/deepmap/oapi-codegen).

```bash
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
oapi-codegen  -package rest openapi.yaml > rest/openapi.gen.go
```

\* For [v0.7.0](https://github.com/unhurried/backend-example-go/tree/v0.7.0) and earlier: [OpenAPI Generator](https://openapi-generator.tech/) was used instead.

```bash
openapi-generator generate -i ./openapi.yaml -g go-gin-server -o ./model --additional-properties=apiPath= --additional-properties=packageName=model
```

### gRPC specification

The specification of gRPC are written in [`todo.proto`](https://github.com/unhurried/backend-example-go/blob/master/grpc/todo.proto).

Source code can be generated with the following command.

```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./grpc/todo.proto
```
