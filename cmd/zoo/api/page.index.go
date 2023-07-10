package api

import (
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

	if err != nil {
		return
	}
	footer.Copyright = CopyrightTag{
		Description: ws.Copyright,
		Link: Link{
			Name: ws.Name,
		},
	}

	footer.Policy, err = getPolicy()
	if err != nil {
		return
	}

	footer.Police, err = getPolice()
	if err != nil {
		return
	}

	footer.Icp, err = getIcp()
	if err != nil {
		return
	}

	footer.CompanyDescribe = CompanyDescribe{
		Logo:     Logo{Link: logoLink},
		Describe: ws.CompanyDescribe,
	}

	footer.CompanyAddress = CompanyAddress{
		Addr:  ws.CompanyAddress,
		Email: ws.CompanyEmail,
		Call:  ws.CompanyTel,
	}

	footer.SiteMap = SiteMap{
		Links: links,
	}

	footer.Links, err = GetFooterLinks()
	if err != nil {
		return
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
