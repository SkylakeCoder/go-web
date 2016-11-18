package web

import "os"

type ENUM_VIEW_TYPE byte

const (
	VIEW_NULL ENUM_VIEW_TYPE = iota
	VIEW_EGO
)

type ViewEngine interface {
	Render(templateRelativePath string, params *ViewParams) (string, error)
}

func getTemplatePath(viewDir string, templateRelativePath string) string {
	if viewDir == "" {
		return templateRelativePath
	}
	return viewDir + string(os.PathSeparator) + templateRelativePath
}
