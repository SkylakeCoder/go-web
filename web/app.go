package web

import (
	"fmt"
	"net/http"
)

// App stands for a application which applies some http services.
type App struct {
	settings *AppSettings
	handler  *appHandler
}

// Get a new App instance.
func NewApp() *App {
	app := &App{
		settings: &AppSettings{},
	}
	app.handler = newAppHandler(app.settings)
	app.registerBuildInPatternHandlers()
	return app
}

// Set the view type.
// see ENUM_VIEW_TYPE in view.go for more details.
func (app *App) SetViewType(viewType ENUM_VIEW_TYPE) error {
	switch viewType {
	case VIEW_EGO:
		app.settings.View = newViewEGO(app.settings)
	default:
		return fmt.Errorf("invalid view type: %v", viewType)
	}
	return nil
}

// Set the view directory.
// You should put all your view templates in the view directory.
func (app *App) SetViewDir(dir string) {
	app.settings.ViewDir = dir
}

// Set the static resource directory.
// Typically, you can put the js, css and img folder in the static resource directory.
func (app *App) SetStaticDir(dir string) {
	app.settings.StaticDir = dir
}

// Handle the http request whose request method is "GET".
func (app *App) Get(pattern string, handler RequestHandler) error {
	err := app.handler.addPatternHandler(HTTP_GET, pattern, handler)
	return err
}

// Handle the http request whose request method is "POST".
func (app *App) Post(pattern string, handler RequestHandler) error {
	err := app.handler.addPatternHandler(HTTP_POST, pattern, handler)
	return err
}

// Register your own PatternHandler.
// See pattern.go for more details.
func (app *App) RegisterPatternHandler(handler PatternHandler) {
	app.handler.registerPatternHandler(handler)
}

// Start to serve.
func (app *App) Listen(port uint32) error {
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.handler)
	return err
}

func (app *App) registerBuildInPatternHandlers() {
	app.RegisterPatternHandler(&patternStatic{})
	app.RegisterPatternHandler(&patternColon{})
}
