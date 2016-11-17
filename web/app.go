package web

import (
	"fmt"
	"log"
	"net/http"
)

type App struct {
	view    ViewEngine
	viewDir string
}

var _appSingleton *App = nil

func GetApp() *App {
	if _appSingleton == nil {
		_appSingleton = &App{}
	}
	return _appSingleton
}

func (app *App) SetViewType(viewType ENUM_VIEW_TYPE) {
	switch viewType {
	case VIEW_EGO:
		app.view = NewViewEGO()
	default:
		log.Fatalln("error view type...")
	}
}

func (app *App) SetViewDir(dir string) {
	app.viewDir = dir
}

func (app *App) Get(pattern string, handler RequestHandler) {
	http.HandleFunc(pattern, func(res http.ResponseWriter, req *http.Request) {
		handler.HandleRequest(
			&Request{
				req,
			},
			&Response{
				res, app.view,
			},
		)
	})
}

func (app *App) Listen(port uint32) {
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatalln(err)
	}
}
