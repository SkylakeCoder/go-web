package web

import (
	"net/http"
	"regexp"
	"strings"
)

type patternColon struct {
	colonReg *regexp.Regexp
}

func (pc *patternColon) HandlePattern(req *http.Request, res http.ResponseWriter, handlersMap HandlersMap, settings *AppSettings) bool {
	// check the situation /xx/:xxx
	if pc.colonReg == nil {
		pc.colonReg = regexp.MustCompile("/:.+")
	}
	patterns := handlersMap.GetPatterns(HTTPMethod(req.Method))
	url := req.URL.Path
	for _, pattern := range patterns {
		reqHandler := handlersMap.GetHandler(HTTPMethod(req.Method), pattern)
		if reqHandler == nil {
			continue
		}
		matches := pc.colonReg.FindAllString(pattern, -1)
		if len(matches) != 1 {
			continue
		}
		urlSplits := strings.Split(url, "/")
		urlSplitsLen := len(urlSplits)
		if urlSplitsLen < 2 {
			continue
		}
		fixedURL := ""
		for i := 0; i < urlSplitsLen-1; i++ {
			fixedURL += urlSplits[i]
			if i < urlSplitsLen-2 {
				fixedURL += "/"
			}
		}
		pattern = strings.Replace(pattern, matches[0], "", -1)
		if pattern == fixedURL {
			paramKey := strings.Replace(matches[0], "/:", "", -1)
			paramValue := urlSplits[len(urlSplits)-1]
			params, _ := NewKeyValues(paramKey, paramValue)
			reqHandler.HandleRequest(
				newRequest(req, params),
				newResponse(res, settings),
			)
			return true
		}
	}
	return false
}
