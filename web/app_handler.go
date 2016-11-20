package web

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type appHandler struct {
	settings      *appSettings
	routerRuleReg *regexp.Regexp
	patternMap    map[string]RequestHandler
}

func newAppHandler(settings *appSettings) *appHandler {
	return &appHandler{
		settings:      settings,
		routerRuleReg: regexp.MustCompile("/:.+"),
		patternMap:    make(map[string]RequestHandler),
	}
}

func (handler *appHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	url := req.URL.Path
	reqHandler, exist := handler.patternMap[url]
	if exist {
		reqHandler.HandleRequest(
			&Request{
				req, nil,
			},
			&Response{
				res,
				handler.settings,
			},
		)
	} else {
		// check the situation /xx/:xxx
		if handler.checkPatternColon(res, req) {
			return
		}
		log.Printf("unrecognized url: %s\n", url)
	}
}

func (handler *appHandler) checkPatternColon(res http.ResponseWriter, req *http.Request) bool {
	url := req.URL.Path
	for k, v := range handler.patternMap {
		matches := handler.routerRuleReg.FindAllString(k, -1)
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
		k = strings.Replace(k, matches[0], "", -1)
		if k == fixedURL {
			paramKey := strings.Replace(matches[0], "/:", "", -1)
			paramValue := urlSplits[len(urlSplits)-1]
			params, _ := NewKeyValues(paramKey, paramValue)
			v.HandleRequest(
				&Request{
					req, params,
				},
				&Response{
					res,
					handler.settings,
				},
			)
			return true
		}
	}
	return false
}

func (handler *appHandler) addPatternHandler(pattern string, reqHandler RequestHandler) error {
	pattern = strings.Replace(pattern, " ", "", -1)
	_, exist := handler.patternMap[pattern]
	if exist {
		return fmt.Errorf("the pattern is exist: %s", pattern)
	}
	handler.patternMap[pattern] = reqHandler
	return nil
}
