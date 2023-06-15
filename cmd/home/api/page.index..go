package api

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/home/internal/db/model"
)

func indexPage(ctx iris.Context) {

	var err error

	var data IndexPage

	ws, err := getWebsite()
	if err != nil {
		return
	}

	data.Metas = buildMetas(ws)

	data.Banner = buildBanner(ws)

	data.Footer, err = buildFooter(ws)
	if err != nil {
		return
	}

	// build header
	data.Header, err = buildHeader(ws)
	if err != nil {
		return
	}

	ctx.ViewLayout("layout/main")
	err = ctx.View("index.html", data)

	if err != nil {
		_, _ = ctx.HTML("<h3>%s</h3>", err.Error())
	}
}

func buildMetas(ws model.WebSite) (metas Metas) {
	metas = Metas{
		Keywords:    ws.PageKeywords,
		Description: ws.PageDescription,
		Title:       ws.PageTitle,
		Author:      ws.PageAuthor,
	}
	return
}

func buildBanner(ws model.WebSite) (banner Banner) {
	banner = Banner{
		Title:    ws.BannerTitle,
		SubTitle: ws.BannerSubTitle,
		Image:    ws.BannerMedia,
	}
	return
}

func buildFooter(ws model.WebSite) (footer Footer, err error) {

	var links []Link
	var logoLink Link

	logoLink, err = getLogo()
	if err != nil {
		return
	}

	copyright, err := getCopyRight()
	if err != nil {
		return
	}
	footer.Copyright = CopyrightTag{
		Description: ws.Copyright,
		Link:        copyright,
	}

	footer.Policy, err = getPolicy()
	if err != nil {
		return
	}
	footer.CompanyDescribe = CompanyDescribe{
		Logo:     Logo{Link: logoLink},
		Describe: ws.CompanyDescribe,
	}

	footer.CompanyAddress = CompanyAddress{
		Addr: ws.CompanyAddress,
		Email: Link{
			Name:  ws.CompanyEmail,
			Outer: true,
		},
		Call: Link{
			Name:  ws.CompanyTel,
			Outer: false,
		},
	}
	footer.CompanyAddress.Email.Url = fmt.Sprintf("mailto:%s", footer.CompanyAddress.Email.Name)
	footer.CompanyAddress.Call.Url = fmt.Sprintf("tel:%s", footer.CompanyAddress.Call.Name)

	friendlyLinks, err := getFriendlyLinks()
	if err != nil {
		return
	}
	footer.FriendlyLinks = FriendlyLinks{
		Links: friendlyLinks,
	}
	links, err = getSiteMap()
	if err != nil {
		return
	}
	footer.SiteMap = SiteMap{
		Links: links,
	}
	return
}

func buildHeader(_ model.WebSite) (header Header, err error) {

	var link Link
	link, err = getLogo()
	if err != nil {
		return
	}
	header.Logo = Logo{Link: link}

	menus, err := getHeaderMenus()
	if err != nil {
		return
	}
	header.Menu.Items = menus
	header.Menu.Login, err = getLogin()
	if err != nil {
		return
	}
	return
}

//func getHeaderMenus() (menus []Link) {
//	menus = []Link{
//		{
//			Url:  "#headere-top",
//			Name: "首页",
//		},
//		{
//			Url:  "#section-one",
//			Name: "关于我们",
//		},
//		{
//			Url:  "#section-two",
//			Name: "产品",
//		},
//		{
//			Url:   "https://qq.com",
//			Name:  "联系方式",
//			Outer: true,
//		},
//		{
//			Url:   "https://baidu.com",
//			Name:  "加入",
//			Outer: true,
//		},
//	}
//	return
//}

//func getLogo() (logo Logo) {
//	logo = Logo{
//		Link: Link{
//			Url:   "https://baidu.com",
//			Media: "assets/images/logo.png",
//		},
//	}
//	return
//}

//func getFriendlyLinks() (links []Link) {
//
//	links = getHeaderMenus()
//	return
//}
