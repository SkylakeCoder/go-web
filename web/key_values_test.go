package web

import (
	"container/list"
	"fmt"
	"testing"
)

func Test_ViewParams(test *testing.T) {
	vp, err := NewKeyValues(
		"key1", 1,
		"key2", "key2value",
		"key3", "key3value",
		"key4", 1.23,
	)
	l := list.New()
	l.PushBack(123)
	l.PushBack("hello")
	l.PushBack(1.23)
	vp.PutList("keyList", l)
	if err != nil {
		test.Fatal("NewKeyValues() failed.")
	}
	keys := vp.GetKeys()
	for _, k := range keys {
		v, err := vp.GetAsString(k)
		if err == nil {
			fmt.Println("key=", k, ", value=", v)
		} else {
			l, _ := vp.GetAsList(k)
			fmt.Print("key= ", k)
			fmt.Print(" , value= [")
			for e := l.Front(); e != nil; e = e.Next() {
				fmt.Print(e.Value, " ")
			}
			fmt.Print("]")
			fmt.Println()
		}
	}
}
