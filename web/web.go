package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type HTTPMethod string

const (
	HTTP_GET    HTTPMethod = "GET"
	HTTP_POST   HTTPMethod = "POST"
	HTTP_PUT    HTTPMethod = "PUT"
	HTTP_DELETE HTTPMethod = "DELETE"
	HTTP_PATCH  HTTPMethod = "PATCH"
)

type Request struct {
	*http.Request
	Params *KeyValues
}

type Response struct {
	http.ResponseWriter
	settings        *AppSettings
	respCache       string
	headerCache     map[string]string
	statusCodeCache int
}

type RequestHandler interface {
	HandleRequest(req *Request, res *Response)
}

func newRequest(req *http.Request, params *KeyValues) *Request {
	newReq := &Request{
		req, params,
	}
	return newReq
}

func newResponse(res http.ResponseWriter, settings *AppSettings) *Response {
	newRes := &Response{
		res, settings,
		"", make(map[string]string),
		http.StatusOK,
	}
	return newRes
}

func (req *Request) GetReqArgs() (url.Values, error) {
	switch HTTPMethod(req.Method) {
	case HTTP_GET:
		fallthrough
	case HTTP_POST:
		return req.Form, nil
	default:
		return nil, fmt.Errorf("the request method [%s] is not supported.", req.Method)
	}
}

func (res *Response) Render(templateRelativePath string, viewParams *KeyValues) (string, error) {
	return res.settings.View.Render(templateRelativePath, viewParams)
}

func (res *Response) SetHeader(key string, value string) {
	res.headerCache[key] = value
}

func (res *Response) SetHeaders(params ...interface{}) error {
	keyValues, err := NewKeyValues(params)
	if err != nil {
		return err
	}
	keys := keyValues.GetKeys()
	for _, k := range keys {
		v, err := keyValues.GetAsString(k)
		if err != nil {
			return err
		}
		res.SetHeader(k, v)
	}
	return nil
}

func (res *Response) SetStatusCode(code int) {
	res.statusCodeCache = code
}

func (res *Response) WriteString(value string) {
	res.respCache += value
}

func (res *Response) WriteJSON(value interface{}) error {
	bytes, err := json.Marshal(value)
	if err == nil {
		res.respCache += string(bytes)
		res.SetHeader("Content-Type", "application/json")
	}
	return err
}

func (res *Response) Flush() error {
	for k, v := range res.headerCache {
		res.Header().Set(k, v)
	}
	res.WriteHeader(res.statusCodeCache)
	_, err := res.Write([]byte(res.respCache))
	return err
}
