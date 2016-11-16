package web

import (
	"testing"
	"log"
)

func Test_ViewParams(test *testing.T) {
	vp, err := NewViewParams(
		"key1", 1,
		"key2", "key2value",
		"key3", "key3value",
		"key4", 1.23,
	)
	if err != nil {
		test.Fatal("NewViewParams() failed.")
	}
	keys := vp.GetKeys()
	for _, k := range keys {
		v, _ := vp.GetAsString(k)
		log.Println("key=", k, ", value=", v)
	}
}
