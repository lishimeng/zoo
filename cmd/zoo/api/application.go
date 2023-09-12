package api

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter/cms"
	"github.com/lishimeng/app-starter/midware/template"
	"github.com/lishimeng/app-starter/midware/web"
	"github.com/lishimeng/home/cmd/zoo/static"
	"github.com/lishimeng/home/internal/etc"
	"time"
)

const AppName = "zoo"

func Application(app *iris.Application) {

	if etc.Config.Web.Cache > 0 { // 设置了页面cache
		app.Use(iris.Cache304(time.Hour * time.Duration(etc.Config.Web.Cache)))
	}

	cms.Init(cms.WithName(AppName), cms.WithDatabase())

	engine := iris.HTML(static.Static, ".html")
	engine.AddFunc("indent", template.Indent)
	engine.AddFunc("nindent", template.NIndent)
	engine.AddFunc("mod", template.NMod)
	app.RegisterView(engine)
	app.Get("/", web.Website, indexPage)
	return
}
