package handlers

import "github.com/SkylakeCoder/go-web/web"

type Hello struct{}

type helloData struct {
	A, B, C int
}

func (h *Hello) HandleRequest(req *web.Request, res *web.Response) {
	res.WriteString("hello go-web !\n")
	data := &helloData{1, 2, 3}
	res.WriteJSON(data)
	res.Flush()
}
