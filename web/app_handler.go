package web

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// HandlersMap is a map that matches the request url and it's handler in a http method.
type HandlersMap map[HTTPMethod]map[string]RequestHandler

// GetPatterns gets the pattern(url) list in a http method.
func (hm HandlersMap) GetPatterns(method HTTPMethod) []string {
	result := []string{}
	subMap, exist := hm[method]
	if exist {
		for k := range subMap {
			result = append(result, k)
		}
	}
	return result
}

// GetHandler gets the handler matches the request method and url.
func (hm HandlersMap) GetHandler(method HTTPMethod, pattern string) RequestHandler {
	subMap, exist := hm[method]
	if !exist {
		return nil
	}
	h, exist := subMap[pattern]
	if !exist {
		return nil
	}
	return h
}

// appHandler is a http.Handler object.
// it finds the correct RequestHandler to handle the http requests.
type appHandler struct {
	settings      *AppSettings
	handlersMap   HandlersMap
	extraPatterns []PatternHandler
}

func newAppHandler(settings *AppSettings) *appHandler {
	return &appHandler{
		settings:    settings,
		handlersMap: make(HandlersMap),
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
		for _, patternHandler := range handler.extraPatterns {
			if patternHandler.HandlePattern(req, res, handler.handlersMap, handler.settings) {
				return
			}
		}
		log.Printf("unrecognized url: %s\n", url)
		handler.handlePattern404(res, req)
	}
}

func (handler *appHandler) registerPatternHandler(patternHandler PatternHandler) {
	handler.extraPatterns = append(handler.extraPatterns, patternHandler)
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
