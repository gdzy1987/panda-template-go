package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/XM-GO/panda-kit/app"
	"github.com/XM-GO/panda-kit/log"
	"github.com/XM-GO/panda-kit/transport"
	"github.com/XM-GO/panda-template-go/pkg/server"
	"github.com/XM-GO/panda-template-go/pkg/service"

	// User import.
	helloworld "github.com/XM-GO/panda-template-go/api/helloworld/v1"
)

var (
	// Name app.
	Name string
	// HTTPAddr string.
	HTTPAddr string
	// GRPCAddr string.
	GRPCAddr string
)

func init() {
	flag.StringVar(&Name, "name", "panda-hello", "app name.")
	flag.StringVar(&HTTPAddr, "http_addr", ":31234", "http listen address.")
	flag.StringVar(&GRPCAddr, "grpc_addr", ":31233", "grpc listen address.")
}

func main() {
	flag.Parse()

	httpSrv := server.NewHTTPServer(HTTPAddr)
	grpcSrv := server.NewGRPCServer(GRPCAddr)
	serverList := []transport.Server{httpSrv, grpcSrv}

	app := app.New(Name,
		&log.Conf{
			App:   Name,
			Level: "debug",
			Dev:   true,
		},
		serverList...,
	)

	{ // User service
		GreeterSrv := service.NewGreeterService()
		helloworld.RegisterGreeterHTTPServer(httpSrv.Container, GreeterSrv)
		helloworld.RegisterGreeterServer(grpcSrv.GetServe(), GreeterSrv)
	}

	if err := app.Run(context.TODO()); err != nil {
		panic(err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, os.Interrupt)
	<-stop

	if err := app.Stop(context.TODO()); err != nil {
		panic(err)
	}
}
