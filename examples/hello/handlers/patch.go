package handlers

import (
	"fmt"
	"github.com/SkylakeCoder/go-web/web"
)

type Patch struct{}

// Usage: curl -X PATCH -d "patchValue=HelloPatch" http://localhost:8688/patch
func (p *Patch) HandleRequest(req *web.Request, res *web.Response) {
	resp := fmt.Sprintf("your method is [%s].\n", req.Method)
	resp += fmt.Sprintf("your patchValue is [%s]\n", req.FormValue("patchValue"))
	res.WriteString(resp)
	res.Flush()
}
