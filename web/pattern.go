package web

import "net/http"

type PatternHandler interface {
	HandlePattern(req *http.Request, res http.ResponseWriter, handlersMap HandlersMap, settings *AppSettings) bool
}
