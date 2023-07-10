package api

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/home/internal/db/model"
	"github.com/lishimeng/home/internal/etc"
)

func indexPage(ctx iris.Context) {

	var err error

	var data IndexPage

	var wsId = etc.Config.Web.WebSiteId
	if wsId <= 0 {
		ctx.StatusCode(tool.RespCodeNotFound)
		return
	}

	ws, err := getWebsite(wsId)
	if err != nil {
		return
	}

	data.Metas = buildMetas(ws)

	data.Banner, err = buildBanner(ws)
	if err != nil {
		return
	}

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
		Favicon:     ws.Favicon,
	}
	return
}

func buildBanner(ws model.WebSite) (banner Banner, err error) {
	banner = Banner{
		Title:    ws.BannerTitle,
		SubTitle: ws.BannerSubTitle,
		Image:    ws.BannerMedia,
	}
	links, err := GetBannerLinks(ws.Id)
	if err != nil {
		return
	}
	banner.Links = links
	return
}

func buildFooter(ws model.WebSite) (footer Footer, err error) {

	var links []Link
	var logoLink Link

	logoLink, err = getLogo(ws.Id)
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

	footer.Policy, err = getPolicy(ws.Id)
	if err != nil {
		return
	}

	footer.Police, err = getPolice(ws.Id)
	if err != nil {
		return
	}

	footer.Icp, err = getIcp(ws.Id)
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

	footer.Links, err = GetFooterLinks(ws.Id)
	if err != nil {
		return
	}
	return
}

func buildHeader(ws model.WebSite) (header Header, err error) {

	var link Link
	link, err = getLogo(ws.Id)
	if err != nil {
		return
	}
	header.Logo = Logo{Link: link}

	menus, err := getHeaderMenus(ws.Id)
	if err != nil {
		return
	}
	header.Menu.Items = menus
	header.Menu.Login, err = getLogin(ws.Id)
	if err != nil {
		return
	}
	return
}
