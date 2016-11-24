package handlers

import (
	"fmt"
	"github.com/SkylakeCoder/go-web/web"
)

type PutForm struct{}

// Usage: curl -X PUT -d "putValue=HelloPut" http://localhost:8688/put_form
func (put *PutForm) HandleRequest(req *web.Request, res *web.Response) {
	resp := fmt.Sprintf("your PUT value is: %s\n", req.FormValue("putValue"))
	res.WriteString(resp)
	res.Flush()
}
