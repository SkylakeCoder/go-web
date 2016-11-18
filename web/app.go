package web

import (
	"fmt"
	"net/http"
)

type App struct {
	settings *appSettings
}

var _appSingleton *App = nil

func GetApp() *App {
	if _appSingleton == nil {
		_appSingleton = &App{
			settings: &appSettings{},
		}
	}
	return _appSingleton
}

func (app *App) SetViewType(viewType ENUM_VIEW_TYPE) error {
	switch viewType {
	case VIEW_EGO:
		app.settings.view = NewViewEGO(app.settings)
	default:
		return fmt.Errorf("invalid view type: %s", viewType)
	}
	return nil
}

func (app *App) SetViewDir(dir string) {
	app.settings.viewDir = dir
}

func (app *App) Get(pattern string, handler RequestHandler) {
	http.HandleFunc(pattern, func(res http.ResponseWriter, req *http.Request) {
		handler.HandleRequest(
			&Request{
				req,
			},
			&Response{
				res,
				app.settings,
			},
		)
	})
}

func (app *App) Listen(port uint32) error {
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	return err
}
