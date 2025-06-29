package singlepage

import (
	"live-server/logger"
	"live-server/watcher"
	"net/http"
	"os"
	"regexp"
)

type httpHandleFunc func(w http.ResponseWriter, r *http.Request)

const fileMatchRegexStr string = `^\/(.*\/)?.*\.[a-zA-Z0-9?_-]+$`

var fileMatchRegexp regexp.Regexp = *regexp.MustCompile(fileMatchRegexStr)

func checkFileUrl(url string) bool {
	return fileMatchRegexp.MatchString(url)
}

func Handler(basePath string, spaEntry string, logEnabled bool, watch bool) httpHandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.String()
		isStaticFile := checkFileUrl(url)

		if isStaticFile {
			logger.Log(logEnabled, "REQUEST: "+url+"  RESPONSE: "+url)
			http.FileServer(http.Dir(basePath)).ServeHTTP(w, r)
		} else {
			indexContent, err := os.ReadFile(basePath + spaEntry)
			if err != nil {
				logger.Log(logEnabled, "REQUEST: "+url+"  RESPONSE: [Failed] not found "+spaEntry)
				http.Error(w, "not found "+spaEntry, http.StatusNotFound)
				return
			}

			if watch {
				w.Write(watcher.IntegrateWatchScript(indexContent))
			} else {
				w.Write(indexContent)
			}
		}
	}
}
