package handlers

import "github.com/SkylakeCoder/go-web/web"

type Hello struct{}

func (h *Hello) HandleRequest(req *web.Request, res *web.Response) {
	res.WriteString("hello go-web !\n")
	res.Flush()
}
