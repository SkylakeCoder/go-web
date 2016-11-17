package web

import "net/http"

type Request struct {
	*http.Request
}

type Response struct {
	http.ResponseWriter
}

type RequestHandler interface {
	HandleRequest(req *Request, res *Response)
}

func (res *Response) Render(templateRelativePath string, viewParams *ViewParams) (string, error) {
	return globalContext.view.Render(templateRelativePath, viewParams)
}

func (res *Response) WriteString(value string) {
	res.Write([]byte(value))
}