package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"
)

const fileMatchRegexStr string = `^\/(.*\/)?.*\.[a-zA-Z0-9?_-]+$`

var fileMatchRegexp regexp.Regexp = *regexp.MustCompile(fileMatchRegexStr)

func checkFileUrl(url string) bool {
	return fileMatchRegexp.MatchString(url)
}

func main() {
	var (
		port string
		src  string
		spa  bool
	)
	const (
		portDefault = "8080"
		portUsage   = "Port to run the server on"
		srcDefault  = ""
		srcUsage    = "Relative path to files"
		spaDefault  = false
		spaUsage    = "Use server for SPA. Server redirects any request to the ./index.html"
	)
	flag.StringVar(&port, "port", portDefault, portUsage)
	flag.StringVar(&port, "p", portDefault, portUsage+" (shorthand)")
	flag.StringVar(&src, "src", srcDefault, srcUsage)
	flag.BoolVar(&spa, "spa", spaDefault, spaUsage)
	flag.Parse()

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting current directory:", err)
	}

	filesPath := currentDir + "/" + src
	if spa {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			url := r.URL.String()
			isStaticFile := checkFileUrl(url)
			log.Println(url, isStaticFile)
			var relativePath string
			if isStaticFile {
				relativePath = url
			} else {
				relativePath = "/index.html"
			}
			f, err := os.Open(filesPath + relativePath)
			if err != nil {
				http.Error(w, "not found "+relativePath, http.StatusNotFound)
			} else {
				http.ServeContent(w, r, relativePath, time.Now(), f)
			}
		})
	} else {
		http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(filesPath))))
	}
	log.Println("Starting live on port:", port, "in path:", filesPath, " ...")
	http.ListenAndServe(":"+port, nil)
}
