package web

import (
	"fmt"
	"net/http"
)

type App struct {
	//empty now.
}

var _appSingleton *App = nil

func GetApp() *App {
	if _appSingleton == nil {
		_appSingleton = &App{}
	}
	return _appSingleton
}

func (app *App) SetViewType(viewType ENUM_VIEW_TYPE) error {
	switch viewType {
	case VIEW_EGO:
		globalContext.view = NewViewEGO()
	default:
		return fmt.Errorf("invalid view type: %s", viewType)
	}
	return nil
}

func (app *App) SetViewDir(dir string) {
	globalContext.viewDir = dir
}

func (app *App) Get(pattern string, handler RequestHandler) {
	http.HandleFunc(pattern, func(res http.ResponseWriter, req *http.Request) {
		handler.HandleRequest(
			&Request{
				req,
			},
			&Response{
				res,
			},
		)
	})
}

func (app *App) Listen(port uint32) error {
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	return err
}
