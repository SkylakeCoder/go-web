package web

import "net/http"

type Request struct {
	*http.Request
	Params *KeyValues
}

type Response struct {
	http.ResponseWriter
	settings *appSettings
}

type RequestHandler interface {
	HandleRequest(req *Request, res *Response)
}

func (res *Response) Render(templateRelativePath string, viewParams *KeyValues) (string, error) {
	return res.settings.view.Render(templateRelativePath, viewParams)
}

func (res *Response) WriteString(value string) {
	res.Write([]byte(value))
}
