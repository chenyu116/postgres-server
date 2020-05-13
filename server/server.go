package server

import (
	"github.com/chenyu116/log"
	"github.com/chenyu116/postgres-server/config"
	"github.com/chenyu116/postgres-server/logger"
	pb "github.com/chenyu116/postgres-server/proto"
	"github.com/chenyu116/postgres-server/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"sync"
)

// Server Server
type Server struct {
	waitGroup *sync.WaitGroup
}

// NewServer NewServer
func NewServer() *Server {
	s := &Server{
		waitGroup: &sync.WaitGroup{},
	}
	return s
}

// Start Start server
func (s *Server) Start() {
	utils.InitPool()

	logger.InitLogger(true, "debug")
	cf := config.GetConfig()

	lis, err := net.Listen("tcp", cf.Serve.GrpcHostPort)
	if err != nil {
		logger.ZapLogger.Fatal("", zap.Error(err))
	}
	grpcServer := grpc.NewServer()
	pb.RegisterApiServer(grpcServer, &apiServer{cfg: cf})
	log.Infof("GRPC Serving at: \"%s\"\n", cf.Serve.GrpcHostPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
