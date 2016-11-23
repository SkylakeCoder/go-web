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

// Request object stands for a http request.
type Request struct {
	*http.Request
	Params *KeyValues
}

// Response object stands for a http response.
type Response struct {
	http.ResponseWriter
	settings        *AppSettings
	respCache       string
	headerCache     map[string]string
	statusCodeCache int
}

// RequestHandler interface helps people to make their own request handler.
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

// GetReqArgs returns the form value if the request method is "GET" or "POST".
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

// Render parses a view template.
func (res *Response) Render(templateRelativePath string, viewParams *KeyValues) (string, error) {
	return res.settings.View.Render(templateRelativePath, viewParams)
}

// SetHeader sets the response header.
func (res *Response) SetHeader(key string, value string) {
	res.headerCache[key] = value
}

// SetHeaders sets the response headers.
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

// SetStatusCode sets the status code of response.
func (res *Response) SetStatusCode(code int) {
	res.statusCodeCache = code
}

// WriteString writes string data to the response cache.
// The response won't be send util you call the Flush() method.
func (res *Response) WriteString(value string) {
	res.respCache += value
}

// WriteJSON writes json data to the response cache.
// The response won't be send util you call the Flush() method.
func (res *Response) WriteJSON(value interface{}) error {
	bytes, err := json.Marshal(value)
	if err == nil {
		res.respCache += string(bytes)
		res.SetHeader("Content-Type", "application/json")
	}
	return err
}

// Flush the response.
func (res *Response) Flush() error {
	for k, v := range res.headerCache {
		res.Header().Set(k, v)
	}
	res.WriteHeader(res.statusCodeCache)
	_, err := res.Write([]byte(res.respCache))
	return err
}
