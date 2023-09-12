package api

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/home/internal/db/model"
)

func indexPage(ctx iris.Context) {

	var err error

	var data IndexPage

	ws, err := getWebsite(AppName)
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

	tool.ResponseHtml(ctx, "layout/main", "index.html", data)

	if err != nil {
		_, _ = ctx.HTML("<h3>%s</h3>", err.Error())
	}
}

func buildMetas(ws model.EnterpriseWebSite) (metas Metas) {
	metas = Metas{
		Keywords:    ws.PageKeywords,
		Description: ws.PageDescription,
		Title:       ws.PageTitle,
		Author:      ws.PageAuthor,
	}
	return
}

func buildBanner(ws model.EnterpriseWebSite) (banner Banner, err error) {
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

func buildFooter(ws model.EnterpriseWebSite) (footer Footer, err error) {

	var links []Link
	var logoLink Link

	logoLink, err = getLogo(ws.Id)
	if err != nil {
		return
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

func buildHeader(ws model.EnterpriseWebSite) (header Header, err error) {

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
