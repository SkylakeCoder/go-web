package handlers

import (
	"container/list"
	"github.com/SkylakeCoder/go-web/web"
)

type View struct{}

func (v *View) HandleRequest(req *web.Request, res *web.Response) {
	params, _ := web.NewKeyValues(
		"title", "go-web",
		"text", "Hello go-web.",
	)
	items := list.New()
	items.PushBack(1)
	items.PushBack("hello world")
	items.PushBack(1.23456)
	params.PutList("items", items)

	result, err := res.Render("test.ego", params)
	if err != nil {
		res.WriteString("error happends when render template: " + err.Error())
	} else {
		res.WriteString(result)
	}
	res.Flush()
}
