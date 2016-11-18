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

func getTemplatePath(relativePath string) string {
	if globalContext.viewDir == "" {
		return relativePath
	}
	return globalContext.viewDir + string(os.PathSeparator) + relativePath
}
