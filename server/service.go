package server

import (
	"context"

	"github.com/atmom/compt/config"
	compt "github.com/atmom/compt/grpc"
	"github.com/sirupsen/logrus"
)

type service struct {
	compt.UnimplementedComptServer
	Config *config.ComptServerConfig
}

// Flow finish request.
func (service *service) FlowFinish(ctx context.Context, request *compt.FlowFinishRequest) (*compt.FlowFinishResponse, error) {
	logrus.Infof("Received: %v", request)
	return &compt.FlowFinishResponse{Code: 0, Message: "success"}, nil
}
