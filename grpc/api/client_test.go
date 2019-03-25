package api

import (
	"os"
	"reflect"
	"testing"
)

var (
	pwd, _    = os.Getwd()
	workDir   = "/tmp/build/"
	volumeDir = pwd + workDir
)

func TestNewGitCloneRequest(t *testing.T) {
	applicationName := "awesome-app"
	jobDir := volumeDir
	repo := "git@git.repo"

	request := NewGitCloneRequest(applicationName, jobDir, repo)
	if request.ApplicationName != "awesome-app" {
		t.Error("application name should match")
	}

	if request.JobDir != volumeDir {
		t.Error("job dir should match")
	}

	if request.Repository != "git@git.repo" {
		t.Error("repository should match")
	}
}

func TestNewExecuteCommandsRequest(t *testing.T) {
	applicationName := "awesome-app"
	jobDir := volumeDir

	request := NewExecuteCommandsRequest(applicationName, jobDir, "echo test", "ls -al")
	if request.ApplicationName != "awesome-app" {
		t.Error("application name should match")
	}

	if request.JobDir != volumeDir {
		t.Error("job dir should match")
	}

	if !reflect.DeepEqual(request.Commands, []string{"echo test", "ls -al"}) {
		t.Error("commands should match")
	}
}
