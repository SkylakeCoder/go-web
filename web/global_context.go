package web

var globalContext context

type context struct {
	view    ViewEngine
	viewDir string
}
