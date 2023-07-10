package api

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter/midware/template"
	"github.com/lishimeng/home/cmd/home/static"
	"github.com/lishimeng/home/internal/etc"
	"time"
)

func Application(app *iris.Application) {

	if etc.Config.Web.Cache > 0 { // 设置了页面cache
		app.Use(iris.Cache304(time.Hour * time.Duration(etc.Config.Web.Cache)))
	}

	engine := iris.HTML(static.Static, ".html")
	engine.AddFunc("indent", template.Indent)
	engine.AddFunc("nindent", template.NIndent)
	engine.AddFunc("mod", template.NMod)
	app.RegisterView(engine)
	app.Get("/", indexPage)
	return
}
