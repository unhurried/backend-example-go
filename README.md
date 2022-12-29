﻿# backend-example-go

A simple Go project that serves REST APIs for TODO web app.

## Covered Features

* API routing, (de)serialization and validation of request/response with [Gin](https://gin-gonic.com/).

## Development Guide

You can run and build the project with `go` command.

```bash
# Run - http server will start and listen on port 3001
go run main.go
# Build - backend(.exe) will be created at the project root directory
go build
```

Also, [`launch.json`](./.vscode/launch.json) for Visual Studio Code is included in the project directory.