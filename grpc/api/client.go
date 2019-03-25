package api

import (
	"context"
	"log"
	"time"

	pb "github.com/joyciapp/joyci-grpc/grpc/proto"
	"google.golang.org/grpc"
)

const address = "localhost:50051"

func connect(address string) *grpc.ClientConn {
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}

func newClient(conn *grpc.ClientConn) pb.JoyciCoreClient {
	return pb.NewJoyciCoreClient(conn)
}

func newContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 15*time.Second)
}

// NewGitCloneRequest initializes a new GitCloneRequest
func NewGitCloneRequest(applicationName string, jobDir string, repository string) *pb.GitCloneRequest {
	return &pb.GitCloneRequest{
		ApplicationName: applicationName,
		JobDir:          jobDir,
		Repository:      repository,
	}
}

// GitClone clones a git repository
func GitClone(applicationName string, jobDir string, repository string) {
	conn := connect(address)
	defer conn.Close()

	c := newClient(conn)

	ctx, cancel := newContext()
	defer cancel()

	if err := ctx.Err(); err != nil {
		log.Fatal("error context:", err)
	}

	request := NewGitCloneRequest(applicationName, jobDir, repository)
	_, err := c.GitClone(ctx, request)
	if err != nil {
		log.Fatal("error on clone a repository:", err)
	}
}

// NewExecuteCommandsRequest initializes a new ExecuteCommandsRequest
func NewExecuteCommandsRequest(applicationName string, jobDir string, commands ...string) *pb.ExecuteCommandsRequest {
	return &pb.ExecuteCommandsRequest{
		ApplicationName: applicationName,
		JobDir:          jobDir,
		Commands:        commands,
	}
}

// ExecuteCommands execute bash commands
func ExecuteCommands(applicationName string, jobDir string, commands ...string) {
	conn := connect(address)
	defer conn.Close()

	c := newClient(conn)

	ctx, cancel := newContext()
	defer cancel()

	if err := ctx.Err(); err != nil {
		log.Fatal("error context:", err)
	}

	request := NewExecuteCommandsRequest(applicationName, jobDir, commands...)
	_, err := c.ExecuteCommands(ctx, request)
	if err != nil {
		log.Fatal("error on execute commands:", err)
	}
}
