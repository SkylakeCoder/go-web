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
	handlersMap   map[HTTPMethod]map[string]RequestHandler
}

func newAppHandler(settings *appSettings) *appHandler {
	return &appHandler{
		settings:      settings,
		routerRuleReg: regexp.MustCompile("/:.+"),
		handlersMap:   make(map[HTTPMethod]map[string]RequestHandler),
	}
}

func (handler *appHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	url := req.URL.Path
	patternMap, exist := handler.handlersMap[HTTPMethod(req.Method)]
	if !exist {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	reqHandler, exist := patternMap[url]
	if exist {
		reqHandler.HandleRequest(
			newRequest(req, nil),
			newResponse(res, handler.settings),
		)
	} else {
		// check if it's a static resource.
		if handler.handlePatternStatic(res, req) {
			return
		}
		// check the situation /xx/:xxx
		if handler.handlePatternColon(res, req) {
			return
		}
		// handle 404...
		log.Printf("unrecognized url: %s\n", url)
		handler.handlePattern404(res, req)
	}
}

func (handler *appHandler) handlePatternStatic(res http.ResponseWriter, req *http.Request) bool {
	url := req.URL.Path
	tmpURL := strings.Trim(url, "/")
	staticDir := strings.Trim(handler.settings.staticDir, "/")
	staticDir = strings.TrimLeft(staticDir, "./")
	if strings.HasPrefix(tmpURL, staticDir) {
		handler := http.FileServer(http.Dir("."))
		handler.ServeHTTP(res, req)
		return true
	}
	return false
}

func (handler *appHandler) handlePatternColon(res http.ResponseWriter, req *http.Request) bool {
	patternMap, exist := handler.handlersMap[HTTPMethod(req.Method)]
	if !exist {
		return false
	}
	url := req.URL.Path
	for k, v := range patternMap {
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
				newRequest(req, params),
				newResponse(res, handler.settings),
			)
			return true
		}
	}
	return false
}

func (handler *appHandler) handlePattern404(res http.ResponseWriter, req *http.Request) bool {
	patternMap, exist := handler.handlersMap[HTTPMethod(req.Method)]
	if !exist {
		return false
	}
	notFoundHandler, exist := patternMap["/404"]
	if exist {
		notFoundHandler.HandleRequest(
			newRequest(req, nil),
			newResponse(res, handler.settings),
		)
		return true
	}
	return false
}

func (handler *appHandler) addPatternHandler(method HTTPMethod, pattern string, reqHandler RequestHandler) error {
	_, exist := handler.handlersMap[method]
	if !exist {
		handler.handlersMap[method] = make(map[string]RequestHandler)
	}
	patternMap, _ := handler.handlersMap[method]
	pattern = strings.Replace(pattern, " ", "", -1)
	_, exist = patternMap[pattern]
	if exist {
		return fmt.Errorf("the pattern is exist: %s", pattern)
	}
	patternMap[pattern] = reqHandler
	return nil
}
