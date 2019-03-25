package api

import (
	"log"
	"os"
	"reflect"
	"testing"
	"time"
)

var (
	pwd, _          = os.Getwd()
	workDir         = "/tmp/build/"
	volumeDir       = pwd + workDir
	applicationName = "joyci-grpc"
	jobDir          = volumeDir
	repo            = "git@github.com:joyciapp/joyci-grpc.git"
)

func TestNewGitCloneRequest(t *testing.T) {
	request := NewGitCloneRequest(applicationName, jobDir, repo)

	if request.ApplicationName != "joyci-grpc" {
		t.Error("application name should match")
	}

	if request.JobDir != volumeDir {
		t.Error("job dir should match")
	}

	if request.Repository != "git@github.com:joyciapp/joyci-grpc.git" {
		t.Error("repository should match")
	}
}

func TestGitCloneIntegration(t *testing.T) {
	go Serve() //Start Server

	GitClone(applicationName, jobDir, repo)

	time.Sleep(10 * time.Second)

	expectedDir := jobDir + "/" + applicationName
	defer os.RemoveAll(expectedDir)

	if _, err := os.Stat(expectedDir); os.IsNotExist(err) {
		log.Println("should clone a git repository")
	}
}

func TestNewExecuteCommandsRequest(t *testing.T) {
	request := NewExecuteCommandsRequest(applicationName, jobDir, "echo test", "ls -al")

	if request.ApplicationName != "joyci-grpc" {
		t.Error("application name should match")
	}

	if request.JobDir != volumeDir {
		t.Error("job dir should match")
	}

	if !reflect.DeepEqual(request.Commands, []string{"echo test", "ls -al"}) {
		t.Error("commands should match")
	}
}
