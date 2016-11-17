package web

type ENUM_VIEW_TYPE byte

const (
	VIEW_NULL ENUM_VIEW_TYPE = iota
	VIEW_EGO
)

type ViewEngine interface {
	Render(templateRelativePath string, params *ViewParams) (string, error)
}