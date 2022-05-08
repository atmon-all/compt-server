package server

import (
	"fmt"
	"net"

	"github.com/atmom/compt/config"
	compt "github.com/atmom/compt/grpc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// start server.
func Start(config *config.ComptServerConfig) {
	if config.Port < 1024 || config.Port > 65536 {
		logrus.Fatalf("Component server start error, invalid server port %d.", config.Port)
	}

	// listen tcp.
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		logrus.Fatalf("Component server start error %v.", err)
	}

	// create server.
	s := grpc.NewServer()
	compt.RegisterComptServer(s, &service{Config: config})

	logrus.Infof("Component server start success, listening at %v", listen.Addr())

	if err := s.Serve(listen); err != nil {
		logrus.Fatalf("Component server start error %v.", err)
	}
}
