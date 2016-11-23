package web

import "net/http"

// PatternHandler interface helps people make their own PatterHandler
// to handle the special url format, such as /user/:username (it's an build-in pattern handler).
type PatternHandler interface {
	HandlePattern(req *http.Request, res http.ResponseWriter, handlersMap HandlersMap, settings *AppSettings) bool
}
