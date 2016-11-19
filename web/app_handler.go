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
	url := req.URL.String()
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
		served := false
		// check the situation /xx/:xxx
		for k, v := range handler.patternMap {
			matches := handler.routerRuleReg.FindAllString(k, -1)
			if len(matches) != 1 {
				continue
			}
			k = strings.Replace(k, matches[0], "", -1)
			if url != k && strings.HasPrefix(url, k) {
				paramKey := strings.Replace(matches[0], "/:", "", -1)
				paramValue := ""
				for i := len(url) - 1; i >= 0; i-- {
					if url[i] == '/' {
						break
					}
					paramValue = string(url[i]) + paramValue
				}
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
				served = true
				break
			}
		}
		if !served {
			log.Printf("unrecognized url: %s\n", url)
		}
	}
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
