package web

import (
	"net/http"
	"strings"
)

type patternStatic struct{}

// HandlePattern handles the static pattern.
func (ps *patternStatic) HandlePattern(req *http.Request, res http.ResponseWriter, handlersMap HandlersMap, settings *AppSettings) bool {
	// check if it's a static resource.
	url := req.URL.Path
	tmpURL := strings.Trim(url, "/")
	staticDir := strings.Trim(settings.StaticDir, "/")
	staticDir = strings.TrimLeft(staticDir, "./")
	if strings.HasPrefix(tmpURL, staticDir) {
		handler := http.FileServer(http.Dir("."))
		handler.ServeHTTP(res, req)
		return true
	}
	return false
}
