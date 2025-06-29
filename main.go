package main

import (
	"flag"
	"fmt"
	"ivnvMkhl/live-server/liveupdate"
	"ivnvMkhl/live-server/logger"
	"ivnvMkhl/live-server/singlepage"
	"net/http"
	"os"
)

var (
	port       string
	src        string
	spaEntry   string
	spa        bool
	logEnabled bool
	watch      bool
)

const (
	spaEntryDefault   = "/index.html"
	spaEntryUsage     = "Path to SPA entry file"
	portDefault       = "8080"
	portUsage         = "Server startup port"
	srcDefault        = ""
	srcUsage          = "Relative path to files"
	spaDefault        = false
	spaUsage          = "Use server for SPA. Server any route request returned ./index.html"
	logEnabledDefault = false
	logEnabledUsage   = "Logging all requests"
	watchDefault      = false
	watchUsage        = "Watch mode for listen modified files in serve path"
)

const liveUpdateWSRoute string = "/ws_live_reload"
const mainRoute string = "/"

func init() {
	flag.StringVar(&port, "port", portDefault, portUsage)
	flag.StringVar(&port, "p", portDefault, portUsage+" (shorthand)")
	flag.StringVar(&src, "src", srcDefault, srcUsage)
	flag.StringVar(&spaEntry, "spa-entry", spaEntryDefault, spaEntryUsage)
	flag.BoolVar(&spa, "spa", spaDefault, spaUsage)
	flag.BoolVar(&logEnabled, "log", logEnabledDefault, logEnabledUsage)
	flag.BoolVar(&watch, "watch", watchDefault, watchUsage)
}

func getWorkingPath() (err error, workingPath string) {
	currentDir, err := os.Getwd()
	if err != nil {
		return err, ""
	}
	return nil, fmt.Sprintf("%s/%s", currentDir, src)
}

func main() {
	flag.Parse()

	err, workingPath := getWorkingPath()
	if err != nil {
		logger.Fatal("Do not get working path ", err)
	}

	if spa {
		http.HandleFunc(mainRoute, singlepage.Handler(workingPath, spaEntry, logEnabled))
	} else {
		http.Handle(mainRoute, http.FileServer(http.Dir(workingPath)))
	}

	if watch {
		http.HandleFunc(liveUpdateWSRoute, liveupdate.Handler(workingPath, logEnabled))
	}

	logger.Log(true, fmt.Sprintf("Starting live on port: %s in path: %s", port, workingPath))
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
