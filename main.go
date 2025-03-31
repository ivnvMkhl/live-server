package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"
)

var (
	port       string
	src        string
	spaEntry   string
	spa        bool
	logEnabled bool
)

const (
	spaEntryDefault   = "/index.html"
	spaEntryUsage     = "Path to SPA entry file. Default: /index.html"
	portDefault       = "8080"
	portUsage         = "Port to run the server on. Default: 8080"
	srcDefault        = ""
	srcUsage          = "Relative path to files"
	spaDefault        = false
	spaUsage          = "Use server for SPA. Server any route request returned ./index.html"
	logEnabledDefault = false
	logEnabledUsage   = "Logging all requests"
)

const fileMatchRegexStr string = `^\/(.*\/)?.*\.[a-zA-Z0-9?_-]+$`

var fileMatchRegexp regexp.Regexp = *regexp.MustCompile(fileMatchRegexStr)

func checkFileUrl(url string) bool {
	return fileMatchRegexp.MatchString(url)
}

func logger(logRestrinction bool, logStr string) {
	if logRestrinction {
		t := time.Now()
		formattedTimestamp := fmt.Sprintf("%d/%02d/%02d %02d:%02d:%02d",
			t.Year(), t.Month(), t.Day(),
			t.Hour(), t.Minute(), t.Second())
		log.Println(formattedTimestamp, " ", logStr)
	}
}

type httpHandleFunc func(w http.ResponseWriter, r *http.Request)

func spaHandler(basePath string, logEnabled bool) httpHandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.String()
		isStaticFile := checkFileUrl(url)

		if isStaticFile {
			logger(logEnabled, "REQUEST: "+url+"  RESPONSE: "+url)
			http.StripPrefix("/", http.FileServer(http.Dir(basePath))).ServeHTTP(w, r)
		} else {
			f, err := os.Open(basePath + spaEntry)
			if err != nil {
				logger(logEnabled, "REQUEST: "+url+"  RESPONSE: [Failed] not found"+spaEntry)
				http.Error(w, "not found "+spaEntry, http.StatusNotFound)
			} else {
				logger(logEnabled, "REQUEST: "+url+"  RESPONSE: "+spaEntry)
				http.ServeContent(w, r, spaEntry, time.Now(), f)
			}
		}
	}
}

func init() {
	flag.StringVar(&port, "port", portDefault, portUsage)
	flag.StringVar(&port, "p", portDefault, portUsage+" (shorthand)")
	flag.StringVar(&src, "src", srcDefault, srcUsage)
	flag.StringVar(&spaEntry, "spa-entry", spaEntryDefault, spaEntryUsage)
	flag.BoolVar(&spa, "spa", spaDefault, spaUsage)
	flag.BoolVar(&logEnabled, "log", logEnabledDefault, logEnabledUsage)
}

func main() {
	flag.Parse()
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting current directory:", err)
	}

	fullPath := fmt.Sprintf("%s/%s", currentDir, src)

	if spa {
		http.HandleFunc("/", spaHandler(fullPath, logEnabled))
	} else {
		http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(fullPath))))
	}

	log.Println("Starting live on port:", port, "in path:", fullPath, " ...")
	http.ListenAndServe(":"+port, nil)
}
