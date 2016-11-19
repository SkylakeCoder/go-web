package web

import (
	"fmt"
	"net/http"
)

type App struct {
	settings *appSettings
	handler  *appHandler
}

var _appSingleton *App = nil

func GetApp() *App {
	if _appSingleton == nil {
		_appSingleton = &App{
			settings: &appSettings{},
		}
		_appSingleton.handler = newAppHandler(_appSingleton.settings)
	}
	return _appSingleton
}

func (app *App) SetViewType(viewType ENUM_VIEW_TYPE) error {
	switch viewType {
	case VIEW_EGO:
		app.settings.view = NewViewEGO(app.settings)
	default:
		return fmt.Errorf("invalid view type: %v", viewType)
	}
	return nil
}

func (app *App) SetViewDir(dir string) {
	app.settings.viewDir = dir
}

func (app *App) Get(pattern string, handler RequestHandler) error {
	err := app.handler.addPatternHandler(pattern, handler)
	return err
}

func (app *App) Listen(port uint32) error {
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.handler)
	return err
}
