package main

import (
	"github.com/SkylakeCoder/go-web/web"
	"handlers"
	"log"
)

func main() {
	app := web.GetApp()
	app.SetViewType(web.VIEW_EGO)
	app.SetViewDir("./views_ego")
	app.SetStaticDir("./static")

	app.Get("/hello", &handlers.Hello{})
	app.Get("/view", &handlers.View{})
	app.Post("/post_form", &handlers.PostForm{})
	app.Get("/user/:username", &handlers.UserColon{})
	app.Get("/count", &handlers.Count{})

	err := app.Listen(8688)
	if err != nil {
		log.Fatalln(err)
	}
}
