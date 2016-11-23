package web

// AppSettings records the current settings.
type AppSettings struct {
	// The ViewEngine you have choose.
	View ViewEngine
	// The view directory you have choose.
	ViewDir string
	// The static resource directory you have choose.
	StaticDir string
}
