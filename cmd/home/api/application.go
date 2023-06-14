package api

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/home/static"
	"html/template"
	"regexp"
	"strings"
)

func Application(app *iris.Application) {

	engine := iris.HTML(static.Static, ".html")
	engine.AddFunc("indent", Indent)
	engine.AddFunc("nindent", nIndent)
	app.RegisterView(engine)
	app.Get("/", index)
	return
}

func Indent(n int, html template.HTML) template.HTML {
	startOfLine := regexp.MustCompile(`(?m)^`)
	indentation := strings.Repeat(" ", n)
	return template.HTML(startOfLine.ReplaceAllLiteralString(string(html), indentation))
}

func nIndent(n int, html template.HTML) template.HTML {
	startOfLine := regexp.MustCompile(`(?m)^`)
	indentation := strings.Repeat(" ", n)
	text := startOfLine.ReplaceAllLiteralString(string(html), indentation)
	return template.HTML(fmt.Sprintf("\n%s", text))
}

func index(ctx iris.Context) {

	var err error

	var data IndexPage

	data.Metas = Metas{
		Keywords:    "主页,cms生成",
		Description: "页面的描述",
		Title:       "一个页面标题,尽量简短",
		Author:      "",
	}

	data.Header.Logo = Logo{
		Url:   "https://baidu.com",
		Image: "assets/images/logo.png",
	}

	data.Banner = Banner{
		Title:    "首页标题",
		SubTitle: "首页子标题，可以写长一点",
		Image:    "assets/images/slider-dec.png",
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
	data.Header.Menu.Login = "https://passport.baidu.com"

	ctx.ViewLayout("layout/main")
	err = ctx.View("index.html", data)

	if err != nil {
		_, _ = ctx.HTML("<h3>%s</h3>", err.Error())
	}
}
