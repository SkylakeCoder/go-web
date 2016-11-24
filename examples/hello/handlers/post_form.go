package handlers

import (
	"fmt"
	"github.com/SkylakeCoder/go-web/web"
)

type PostForm struct{}

func (pf *PostForm) HandleRequest(req *web.Request, res *web.Response) {
	resp := fmt.Sprintf("your POST input is: %s\n", req.FormValue("inputText"))
	res.WriteString(resp)
	res.Flush()
}
