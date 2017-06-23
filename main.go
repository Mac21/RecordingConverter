package main

import (
	"flag"
	"os"

	"path/filepath"

	"time"

	"strings"

	"github.com/golang/glog"
)

var (
	ffmpeg                *FFmpeg
	runningDir            string
	daemon                bool
	sleepDurationMultiple int
)

func init() {
	flag.StringVar(&runningDir, "p", "", "Full path of running directory without trailing slash.")
	flag.IntVar(&sleepDurationMultiple, "s", 0, "The multiple of seconds to sleep before walking running dir again.")
	flag.BoolVar(&daemon, "d", false, "Run as daemon")

	wd, err := os.Getwd()
	if err != nil {
		glog.Fatal(err.Error())
	}

	//C:\Users\Control\go\src\github.com\mac21\RecordingsConverter\ffmpeg\bin
	ffmpeg, err = NewFFmpeg(wd+"\\ffmpeg\\bin\\ffmpeg.exe", "")
	if err != nil {
		glog.Fatal(err.Error())
	}
}

func fileNewExtension(filename, desiredExt string) string {
	parts := strings.Split(filename, ".")
	return strings.Join([]string{parts[0], desiredExt}, ".")
}

func walkCallBack(inFile string, info os.FileInfo, err error) error {
	glog.V(2).Info(inFile)
	if info.IsDir() {
		return nil
	}

	outFileName := fileNewExtension(inFile, "mp3")
	glog.V(2).Info(outFileName)

	ffmpeg.SetArgs("-n", "-i", inFile, outFileName)
	ffmpeg.Run()

	return nil
}

func main() {
	flag.Parse()

	if runningDir == "" {
		glog.Fatal("Running directory flag required!")
	}

	for {
		filepath.Walk(runningDir, walkCallBack)

		if daemon == false {
			break
		}

		glog.V(1).Infof("Finished converting current files. Sleeping for %d seconds.", sleepDurationMultiple)
		time.Sleep(time.Duration(sleepDurationMultiple) * time.Second)
	}

}
