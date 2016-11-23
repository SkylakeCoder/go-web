package web

import "os"

type ENUM_VIEW_TYPE byte

const (
	// No need to use view.
	VIEW_NULL ENUM_VIEW_TYPE = iota
	// VIEW_EGO is a view template that like the ejs in node.js Express framework.
	VIEW_EGO
)

// ViewEngine interface helps people make their own view template parser.
type ViewEngine interface {
	Render(templateRelativePath string, params *KeyValues) (string, error)
}

// A common method to get the template path.
func getTemplatePath(viewDir string, templateRelativePath string) string {
	if viewDir == "" {
		return templateRelativePath
	}
	return viewDir + string(os.PathSeparator) + templateRelativePath
}
