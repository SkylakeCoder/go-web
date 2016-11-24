package handlers

import (
	"fmt"
	"github.com/SkylakeCoder/go-web/web"
	"io/ioutil"
)

type PostUpload struct{}

func (pu *PostUpload) HandleRequest(req *web.Request, res *web.Response) {
	f, header, err := req.FormFile("file")
	if err != nil {
		res.WriteString("upload error: " + err.Error())
		res.Flush()
		return
	}
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		res.WriteString("read file error: " + err.Error())
		res.Flush()
		return
	}
	err = ioutil.WriteFile("./upload_dir/"+header.Filename, bytes, 0644)
	if err != nil {
		res.WriteString("save file failed: " + err.Error())
	} else {
		res.WriteString(fmt.Sprintf("the file %s upload success.", header.Filename))
	}
	res.Flush()
}
