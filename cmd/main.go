package main

import (
	"context"

	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	apipb "github.com/guoxf/dubbogo-demo/api"
)

type HelloWorldProvider struct {
	apipb.UnimplementedHelloWorldServer
}

func (h *HelloWorldProvider) SayHello(ctx context.Context, in *apipb.SayHelloRequest) (*apipb.SayHelloResponse, error) {
	return &apipb.SayHelloResponse{}, nil
}

// gin-swagger middleware
// swagger embed files
func main() {
	config.SetProviderService(&HelloWorldProvider{})
	if err := config.Load(); err != nil {
		panic(err)
	}
	select {}
}
