# go-web
Just a toy web framework written in golang.

## install
```
$ go get github.com/SkylakeCoder/go-web
```

## usage
See more in [examples](https://github.com/SkylakeCoder/go-web/tree/master/examples/hello)
```
package main

import (
	"./handlers"
	"github.com/SkylakeCoder/go-web/web"
	"log"
)

func main() {
	app := web.GetApp()
	app.SetViewType(web.VIEW_EGO)
	app.SetViewDir("./views_ego")
	app.SetStaticDir("./static")

	app.Get("/view", &handlers.View{})
	app.Post("/post_form", &handlers.PostForm{})
	app.Get("/user/:username", &handlers.UserColon{})
	app.Get("/404", &handlers.Handler404{})
	app.Post("/404", &handlers.Handler404{})

	err := app.Listen(8688)
	if err != nil {
		log.Fatalln(err)
	}
}

```
