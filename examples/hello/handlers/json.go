package handlers

import "github.com/SkylakeCoder/go-web/web"

type JSON struct {
	A, B, C int
	Slice   []string
	Map     map[string][]int
}

func (j *JSON) HandleRequest(req *web.Request, res *web.Response) {
	j.A = 1
	j.B = 2
	j.C = 3
	j.Slice = []string{"hello", "world", "golang"}
	j.Map = make(map[string][]int)
	j.Map["key1"] = []int{1, 2, 3}
	j.Map["key2"] = []int{4, 5, 6}
	j.Map["key3"] = []int{7, 8, 9}

	res.WriteJSON(j)
	res.Flush()
}
