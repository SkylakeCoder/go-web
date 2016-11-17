package web

import (
	"fmt"
	"sync"
	"testing"
)

type testHandler struct {
	visitCount uint32
	lock       sync.Mutex
}
type helloHandler struct{}

func (test *testHandler) HandleRequest(req *Request, res *Response) {
	test.lock.Lock()
	test.visitCount++
	defer test.lock.Unlock()

	res.WriteString(fmt.Sprintf("You have visit this page %d times.", test.visitCount))
}

func (hello helloHandler) HandleRequest(req *Request, res *Response) {
	params, _ := NewViewParams(
		"title", "go-web",
		"text", "This is the response from the helloHandler.",
	)
	result, err := res.Render("test.ego", params)
	if err != nil {
		res.WriteString("error happends when render template." + err.Error())
	} else {
		res.WriteString(result)
	}
}

func Test_App(test *testing.T) {
	app := GetApp()
	app.SetViewType(VIEW_EGO)
	app.SetViewDir("./views_ego")
	app.Get("/test", &testHandler{})
	app.Get("/hello", helloHandler{})
	app.Listen(8686)
}
