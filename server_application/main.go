package main

import (
	"flag"
	"github.com/jjmeg/WorkChopProject/app"
	"github.com/jjmeg/WorkChopProject/util/runmode"
	"os"
	"path"
)

var (
	runMode string
	srcPath string
)

func init() {
	flag.StringVar(&runMode, "runMode", "development", "development|test|production")
	flag.StringVar(&srcPath, "srcPath", "", "path")
}

func main() {
	flag.Parse()

	mode := runmode.RunMode(runMode)
	if !mode.IsValid() {
		panic("mode error")
	}

	if srcPath == "" {
		var err error
		srcPath, err = os.Getwd()
		if err != nil {
			panic(err)
		}
	} else {
		srcPath = path.Clean(srcPath)
	}

	app.New(mode, srcPath).Run()
}
