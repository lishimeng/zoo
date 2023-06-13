package api

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/home/static"
)

func Application(app *iris.Application) {
	app.RegisterView(iris.HTML(static.Static, ".html"))
	app.Get("/", index)
	return
}

func index(ctx iris.Context) {

	var err error

	var data IndexPage

	data.Header.Logo = Logo{
		Title: "标题",
		Url:   "#headere-top",
		Image: "images/logo.png",
	}

	data.Header.Banner = Banner{
		Title:    "首页标题",
		SubTitle: "首页子标题，可以写长一点",
		Image:    "images/banner-image.jpg",
	}

	data.Footer.Copyright = CopyrightTag{
		Description: "2023 XXXX. All rights reserved.",
		Name:        "XXXX",
		Link:        "https://xxxx.com",
	}

	data.Footer.Policy = PolicyTag{
		TermOfUse:     "使用条款",
		PrivatePolicy: "隐私政策",
	}

	var menus = []MenuItem{
		{
			Link:   "#headere-top",
			Name:   "首页",
			Active: true,
		},
		{
			Link: "#section-one",
			Name: "关于我们",
		},
		{
			Link: "#section-two",
			Name: "产品",
		},
		{
			Link:  "https://qq.com",
			Name:  "联系方式",
			Outer: true,
		},
		{
			Link:  "https://baidu.com",
			Name:  "加入",
			Outer: true,
		},
	}
	data.Header.Menu.Items = menus

	ctx.ViewLayout("layout/main")
	err = ctx.View("index.html", data)

	if err != nil {
		_, _ = ctx.HTML("<h3>%s</h3>", err.Error())
	}
}
