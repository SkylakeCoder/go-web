package handlers

import (
	"github.com/SkylakeCoder/go-web/web"
	"fmt"
)

type PostForm struct{}

func (pf *PostForm) HandleRequest(req *web.Request, res *web.Response) {
	resp := fmt.Sprintf("your input is: %s", req.FormValue("inputText"))
	res.WriteString(resp)
	res.Flush()
}