package main

import (
	"regexp"
	"testing"
)

func TestCreateFileRegex(t *testing.T) {
	_, err := regexp.Compile(fileMatchRegexStr)

	if err != nil {
		t.Errorf("Failed to create fileMatchRegex " + fileMatchRegexStr)
	}
}

var (
	fileUrlMocks = [9]string{
		"/styles/supportedBrowser.css?1743425377764",
		"/scripts/bowserBundled.js?1743425377764",
		"/scripts/supportedBrowser.js?1743425377764",
		"/scripts/handleAppError.js?1743425377764",
		"/scripts/envConfig.js?1743425377764",
		"/static/js/monaco.dbf45967.chunk.js",
		"/static/js/4188.373ce626.chunk.js",
		"/static/js/node_modules.373ce626.chunk.js",
		"/static/js/app-menu.373ce626.chunk.js",
	}
	routeUrlMocks = [11]string{
		"/static/js/",
		"/static/js",
		"/reports?filter=public",
		"/",
		"/?fullscreen=true",
		"/reports",
		"/report-list",
		"/slice_list",
		"/reports/",
		"/report-list/",
		"/slice_list/",
	}
)

func TestMatchFileRegex(t *testing.T) {
	for _, url := range fileUrlMocks {
		valid := checkFileUrl(url)
		if !valid {
			t.Errorf("File url %s is not match to files regex %s", url, fileMatchRegexStr)
		}
	}

	for _, url := range routeUrlMocks {
		valid := checkFileUrl(url)
		if valid {
			t.Errorf("Route url %s in match to files regex %s", url, fileMatchRegexStr)
		}
	}
}
