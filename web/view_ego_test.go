package web

import (
	"container/list"
	"log"
	"testing"
)

func Test_ViewEGO(test *testing.T) {
	vp, _ := NewViewParams(
		"title", "ego-test",
		"text", "test...",
	)
	items := list.New()
	items.PushBack(1)
	items.PushBack("hello world")
	items.PushBack(1.23456)
	vp.PutList("items", items)

	settings := &appSettings{
		viewDir: "./views_ego",
	}
	ego := NewViewEGO(settings)
	v, err := ego.Render("test.ego", vp)
	if err != nil {
		test.Fatal("Parse error: " + err.Error())
	}
	log.Println(v)
}
