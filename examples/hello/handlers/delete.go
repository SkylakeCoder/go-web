package handlers

import (
	"fmt"
	"github.com/SkylakeCoder/go-web/web"
)

type Delete struct{}

// Usage: curl -X DELETE http://localhost:8688/delete
func (d *Delete) HandleRequest(req *web.Request, res *web.Response) {
	resp := fmt.Sprintf("your method is [%s].\n", req.Method)
	res.WriteString(resp)
	res.Flush()
}
