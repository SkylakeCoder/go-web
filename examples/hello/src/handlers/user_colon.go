package handlers

import "github.com/SkylakeCoder/go-web/web"

type UserColon struct{}

func (user *UserColon) HandleRequest(req *web.Request, res *web.Response) {
	// /user/:username
	userName, _ := req.Params.GetAsString("username")
	res.WriteString("your name is: " + userName)
	res.Flush()
}
