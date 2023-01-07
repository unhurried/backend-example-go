# backend-example-go

A simple Go project that serves REST APIs for TODO web app.

## Covered Features

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

Model and router code can be generated with OpenAPI Generator.

```bash
openapi-generator generate -i ./openapi.yaml -g go-gin-server -o ./model --additional-properties=apiPath= --additional-properties=packageName=model
```
