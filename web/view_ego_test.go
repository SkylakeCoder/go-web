package web

import (
	"log"
	"testing"
)

func Test_ViewEGO(test *testing.T) {
	vp, _ := NewViewParams(
		"title", "ego-test",
		"text", "test...",
	)
	ego := NewViewEGO()
	v, err := ego.Render("test.ego", vp)
	if err != nil {
		test.Fatal("Parse error...")
	}
	log.Println(v)
}
