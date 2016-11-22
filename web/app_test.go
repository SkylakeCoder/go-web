package web

import (
	"container/list"
	"fmt"
	"sync"
	"testing"
)

type testHandler struct {
	visitCount uint32
	lock       sync.Mutex
}
type helloHandler struct{}
type userHandler struct{}
type notFoundHandler struct{}

func (test *testHandler) HandleRequest(req *Request, res *Response) {
	test.lock.Lock()
	test.visitCount++
	defer test.lock.Unlock()

	res.WriteString(fmt.Sprintf("You have visit this page %d times.", test.visitCount))
	res.Flush()
}

func (hello *helloHandler) HandleRequest(req *Request, res *Response) {
	params, _ := NewKeyValues(
		"title", "go-web",
		"text", "This is the response from the helloHandler.",
	)
	items := list.New()
	items.PushBack(1)
	items.PushBack("hello world")
	items.PushBack(1.23456)
	params.PutList("items", items)

	result, err := res.Render("test.ego", params)
	if err != nil {
		res.WriteString("error happends when render template." + err.Error())
	} else {
		res.WriteString(result)
	}
	res.WriteString("key1=" + req.FormValue("key1"))
	res.Flush()
}

func (user *userHandler) HandleRequest(req *Request, res *Response) {
	name, _ := req.Params.GetAsString("username")
	res.WriteString(fmt.Sprintf("your name is: %s", name))
	res.Flush()
}

func (notFound *notFoundHandler) HandleRequest(req *Request, res *Response) {
	res.WriteString(fmt.Sprintf("couldn't find the url: %s", req.URL.Path))
	res.Flush()
}

func Test_App(test *testing.T) {
	app := GetApp()
	app.SetViewType(VIEW_EGO)
	app.SetViewDir("./views_ego")
	app.SetStaticDir("./static")
	app.Get("/test", &testHandler{})
	app.Get("/hello", &helloHandler{})
	app.Post("/hello", &helloHandler{})
	app.Get("/user/:username", &userHandler{})
	app.Get("/404", &notFoundHandler{})
	err := app.Listen(8688)
	if err != nil {
		test.Fatal("app.Listen error: ", err)
	}
}
