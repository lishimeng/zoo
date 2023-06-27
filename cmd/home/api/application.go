package api

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/home/cmd/home/midware"
	"github.com/lishimeng/home/internal/etc"
	"github.com/lishimeng/home/static"
	"time"
)

func Application(app *iris.Application) {

	if etc.Config.Web.Cache > 0 { // 设置了页面cache
		app.Use(iris.Cache304(time.Hour * time.Duration(etc.Config.Web.Cache)))
	}

	engine := iris.HTML(static.Static, ".html")
	engine.AddFunc("indent", midware.Indent)
	engine.AddFunc("nindent", midware.NIndent)
	engine.AddFunc("mod", midware.Mod)
	app.RegisterView(engine)
	app.Get("/", indexPage)
	return
}
