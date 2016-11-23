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
	app.Get("/user/:username", &handlers.UserColon{})
	app.Get("/count", &handlers.Count{})
	app.Get("/404", &handlers.Handler404{})
	app.Post("/404", &handlers.Handler404{})

	err := app.Listen(8688)
	if err != nil {
		log.Fatalln(err)
	}
}
