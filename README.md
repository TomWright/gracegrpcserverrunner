[![Go Report Card](https://goreportcard.com/badge/github.com/TomWright/gracegrpcserverrunner)](https://goreportcard.com/report/github.com/TomWright/gracegrpcserverrunner)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/tomwright/gracegrpcserverrunner)](https://pkg.go.dev/github.com/tomwright/gracegrpcserverrunner)
![GitHub License](https://img.shields.io/github/license/TomWright/gracegrpcserverrunner)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/TomWright/gracegrpcserverrunner?label=latest%20release)

# Grace GRPC Server Runner

A GRPC Server Runner for use with [grace](https://github.com/TomWright/grace).

## Usage

```go
package main

import (
	"context"
	"github.com/tomwright/grace"
	"github.com/tomwright/gracegrpcserverrunner"
	"google.golang.org/grpc"
)

func main() {
	g := grace.Init(context.Background())

	// Create and configure your GRPC server.
	server := grpc.NewServer()

	// Create and configure the GRPC server runner.
	runner := &gracegrpcserverrunner.GRPCServerRunner{
		Server:        server,
		ListenAddress: ":9090",
	}

	// Run the runner.
	g.Run(runner)

	g.Wait()
}
```
