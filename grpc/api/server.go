package api

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/joyciapp/joyci-core/cmd/bash"
	"github.com/joyciapp/joyci-core/cmd/git"
	pb "github.com/joyciapp/joyci-grpc/grpc/proto"
	"google.golang.org/grpc"
)

const (
	defaultPort      = ":50051"
	serverPortEnvKey = "SERVER_PORT"
)

// Server structs representing the GRPC Api server
type Server struct{}

// GitClone implementation
func (s *Server) GitClone(ctx context.Context, request *pb.GitCloneRequest) (*empty.Empty, error) {
	git := git.New().VolumeDir(request.JobDir).Build()
	git.Clone(request.Repository)

	return new(empty.Empty), nil
}

// ExecuteCommands implementation
func (s *Server) ExecuteCommands(ctx context.Context, request *pb.ExecuteCommandsRequest) (*empty.Empty, error) {
	volumeAndWorkDir := request.JobDir + "/" + request.ApplicationName
	bash := bash.New().VolumeAndWorkDir(volumeAndWorkDir).Build()
	bash.Execute(request.Commands...)

	return new(empty.Empty), nil
}

// GetServerListenPort fetch port to be used on server start up
func GetServerListenPort() string {
	if port := os.Getenv(serverPortEnvKey); port != "" {
		return fmt.Sprintf(":%s", port)
	}

	return defaultPort
}

// Serve start grpc server
func Serve() {
	listenPort := GetServerListenPort()
	lis, err := net.Listen("tcp", listenPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterJoyciCoreServer(s, &Server{})

	log.Println("JoyCI GRPC server started at ", listenPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
