package web

import (
	"fmt"
	"net/http"
)

type App struct {
	settings *AppSettings
	handler  *appHandler
}

func GetApp() *App {
	app := &App{
		settings: &AppSettings{},
	}
	app.handler = newAppHandler(app.settings)
	app.registerBuildInPatternHandlers()
	return app
}

func (app *App) SetViewType(viewType ENUM_VIEW_TYPE) error {
	switch viewType {
	case VIEW_EGO:
		app.settings.View = NewViewEGO(app.settings)
	default:
		return fmt.Errorf("invalid view type: %v", viewType)
	}
	return nil
}

func (app *App) SetViewDir(dir string) {
	app.settings.ViewDir = dir
}

func (app *App) SetStaticDir(dir string) {
	app.settings.StaticDir = dir
}

func (app *App) Get(pattern string, handler RequestHandler) error {
	err := app.handler.addPatternHandler(HTTP_GET, pattern, handler)
	return err
}

func (app *App) Post(pattern string, handler RequestHandler) error {
	err := app.handler.addPatternHandler(HTTP_POST, pattern, handler)
	return err
}

func (app *App) RegisterPatternHandler(handler PatternHandler) {
	app.handler.registerPatternHandler(handler)
}

func (app *App) Listen(port uint32) error {
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.handler)
	return err
}

func (app *App) registerBuildInPatternHandlers() {
	app.RegisterPatternHandler(&patternStatic{})
	app.RegisterPatternHandler(&patternColon{})
}
