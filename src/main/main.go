package main

import (
	"flag"
	"fmt"
	"app"
	"os"
)

func parseArgs()(*app.Config, error){
	syslog := flag.Bool("syslog", false, "Log to syslog")

	flag.Parse()

	config := app.NewConfig()
	config.Syslog = *syslog
	return config, nil
}

func main(){
	appName := "myapp"
	logTag := "myapp log"
	config, err := parseArgs()
	var logger *app.Logger
	if config.Syslog {
		logger, err = app.NewSysLogger(logTag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to initialize syslog: %v", err)
			os.Exit(1)
		}
	}else {
		logger = app.NewTerminaLogger(logTag)
	}

	// log test
	logger.Infof("app %v is runing", appName)
	logger.Warnf("There seems to be a warning: %v", "test warning")
	logger.Errorf("func A failed: could not get a status: %v", "test error")
}





