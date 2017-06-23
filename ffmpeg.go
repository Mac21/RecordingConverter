package main

import (
	"os/exec"

	"github.com/golang/glog"
)

// FFmpeg struct is to be constructed with an optional path to the ffmpeg executable.
type FFmpeg struct {
	exeName string
	exePath string
	args    []string
}

// NewFFmpeg creates a new ffmpeg object.
// exePath if not provided defaults to the exe in your PATH.
func NewFFmpeg(exePath string, args ...string) (*FFmpeg, error) {
	exeName := "ffmpeg"

	if exePath == "" {
		exePath, err := exec.LookPath(exeName)
		if err != nil {
			return nil, err
		}
		glog.V(2).Info(exePath)
	}

	return &FFmpeg{exeName, exePath, args}, nil
}

// Run FFmpeg with arguements that were provided to the constructor.
func (f *FFmpeg) Run() {
	cmd := exec.Command(f.exePath, f.args...)

	out, err := cmd.CombinedOutput()
	glog.V(2).Infof("%s\n", out)
	if err != nil {
		glog.V(1).Info(err.Error())
	}
}

// SetArgs overwrite args used when object was constructed.
func (f *FFmpeg) SetArgs(args ...string) {
	f.args = args
}

// GetArgs get current args.
func (f *FFmpeg) GetArgs() []string {
	return f.args
}
