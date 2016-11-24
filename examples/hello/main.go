package main

import (
	"./handlers"
	"github.com/SkylakeCoder/go-web/web"
	"log"
)

func main() {
	app := web.NewApp()
	app.SetViewType(web.VIEW_EGO)
	app.SetViewDir("./views_ego")
	app.SetStaticDir("./static")

	app.Get("/hello", &handlers.Hello{})
	app.Get("/json", &handlers.JSON{})
	app.Get("/view", &handlers.View{})
	app.Post("/post_form", &handlers.PostForm{})
	app.Post("/upload", &handlers.PostUpload{})

	// Usage: curl -X PUT -d "putValue=HelloPut" http://localhost:8688/put_form
	app.Put("/put_form", &handlers.PutForm{})
	// Usage: curl -X DELETE http://localhost:8688/delete
	app.Delete("/delete", &handlers.Delete{})
	// Usage: curl -X PATCH -d "patchValue=HelloPatch" http://localhost:8688/patch
	app.Patch("/patch", &handlers.Patch{})

	app.Get("/user/:username", &handlers.UserColon{})
	app.Get("/count", &handlers.Count{})
	app.Get("/404", &handlers.Handler404{})
	app.Post("/404", &handlers.Handler404{})

	err := app.Listen(8688)
	if err != nil {
		log.Fatalln(err)
	}
}
