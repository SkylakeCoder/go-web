package web

import (
	"fmt"
	"net/http"
	"strings"
)

type patternStatic struct{}

// HandlePattern handles the static pattern.
func (ps *patternStatic) HandlePattern(req *http.Request, res http.ResponseWriter, handlersMap HandlersMap, settings *AppSettings) (bool, bool) {
	// check if it's a static resource.
	url := req.URL.Path
	if strings.HasSuffix(url, "/") {
		res.WriteHeader(http.StatusForbidden)
		resp := fmt.Sprintf("FORBIDDEN: you have no permission to access the path %s\n", url)
		res.Write([]byte(resp))
		return false, false
	}
	tmpURL := strings.Trim(url, "/")
	staticDir := strings.Trim(settings.StaticDir, "/")
	staticDir = strings.TrimLeft(staticDir, "./")
	if strings.HasPrefix(tmpURL, staticDir) {
		handler := http.FileServer(http.Dir("."))
		handler.ServeHTTP(res, req)
		return true, false
	}
	return false, true
}
