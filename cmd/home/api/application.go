package api

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/home/cmd/home/midware"
	"github.com/lishimeng/home/static"
	"time"
)

func Application(app *iris.Application) {

	app.Use(iris.Cache304(time.Second * 10))

	engine := iris.HTML(static.Static, ".html")
	engine.AddFunc("indent", midware.Indent)
	engine.AddFunc("nindent", midware.NIndent)
	engine.AddFunc("mod", midware.Mod)
	app.RegisterView(engine)
	app.Get("/", indexPage)
	return
}
