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

// NewApp returns a new App instance.
func NewApp() *App {
	app := &App{
		settings: &AppSettings{},
	}
	app.handler = newAppHandler(app.settings)
	app.registerBuildInPatternHandlers()
	return app
}

// SetViewType sets the view type.
// See ENUM_VIEW_TYPE in view.go for more details.
func (app *App) SetViewType(viewType ENUM_VIEW_TYPE) error {
	switch viewType {
	case VIEW_EGO:
		app.settings.View = newViewEGO(app.settings)
	default:
		return fmt.Errorf("invalid view type: %v", viewType)
	}
	return nil
}

// SetViewDir sets the view directory.
// You should put all your view templates in the view directory.
func (app *App) SetViewDir(dir string) {
	app.settings.ViewDir = dir
}

// SetStaticDir sets the static resource directory.
// Typically, you can put the js, css and img folder in the static resource directory.
func (app *App) SetStaticDir(dir string) {
	app.settings.StaticDir = dir
}

// Get handles the http request whose request method is "GET".
func (app *App) Get(pattern string, handler RequestHandler) error {
	err := app.handler.addPatternHandler(HTTP_GET, pattern, handler)
	return err
}

// Post handles the http request whose request method is "POST".
func (app *App) Post(pattern string, handler RequestHandler) error {
	err := app.handler.addPatternHandler(HTTP_POST, pattern, handler)
	return err
}

// Put handles the http request whose request method is "PUT".
func (app *App) Put(pattern string, handler RequestHandler) error {
	err := app.handler.addPatternHandler(HTTP_PUT, pattern, handler)
	return err
}

// Patch handles the http request whose request method is "PATCH".
func (app *App) Patch(pattern string, handler RequestHandler) error {
	err := app.handler.addPatternHandler(HTTP_PATCH, pattern, handler)
	return err
}

// Delete handles the http request whose request method is "DELETE".
func (app *App) Delete(pattern string, handler RequestHandler) error {
	err := app.handler.addPatternHandler(HTTP_DELETE, pattern, handler)
	return err
}

// RegisterPatternHandler registers your own PatternHandler.
// See pattern.go for more details.
func (app *App) RegisterPatternHandler(handler PatternHandler) {
	app.handler.registerPatternHandler(handler)
}

// Listen some port and start to serve.
func (app *App) Listen(port uint32) error {
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.handler)
	return err
}

func (app *App) registerBuildInPatternHandlers() {
	app.RegisterPatternHandler(&patternStatic{})
	app.RegisterPatternHandler(&patternColon{})
}
