package web

import (
	"container/list"
	"log"
	"testing"
)

func Test_ViewEGO(test *testing.T) {
	globalContext.viewDir = "./views_ego"
	vp, _ := NewViewParams(
		"title", "ego-test",
		"text", "test...",
	)
	items := list.New()
	items.PushBack(1)
	items.PushBack("hello world")
	items.PushBack(1.23456)
	vp.PutList("items", items)

	ego := NewViewEGO()
	v, err := ego.Render("test.ego", vp)
	if err != nil {
		test.Fatal("Parse error: " + err.Error())
	}
	log.Println(v)
}
