package handlers

import (
	"github.com/SkylakeCoder/go-web/web"
	"fmt"
)

type Handler404 struct {}

func (h404 *Handler404) HandleRequest(req *web.Request, res *web.Response) {
	res.WriteString(fmt.Sprintf("Couldn't find the URL: %s", req.URL.Path))
	res.Flush()
}
