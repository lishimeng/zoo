package api

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/home/internal/db/model"
)

func getWebsite() (ws model.WebSite, err error) {
	var webSite = 1
	var tmp []model.WebSite
	_, err = app.GetOrm().Context.QueryTable(new(model.WebSite)).
		Filter("id", webSite).
		All(&tmp)
	if len(tmp) > 0 {
		ws = tmp[0]
	}

	return
}

func getHeaderMenus() (menus []Link, err error) {
	var webSite = 1
	menus, err = getLinks(webSite, model.CategoryHeaderMenu)
	return
}

func getLogin() (link Link, err error) {
	var webSite = 1
	link, err = getLink(webSite, model.CategoryLogIn)
	return
}

func getLogo() (link Link, err error) {
	var webSite = 1
	link, err = getLink(webSite, model.CategoryLogo)
	return
}

func getFriendlyLinks() (links []Link, err error) {
	var webSite = 1
	links, err = getLinks(webSite, model.CategoryFriendlyLink)
	return
}

func getCopyRight() (link Link, err error) {
	var webSite = 1
	link, err = getLink(webSite, model.CategoryCopyright)
	return
}

func getPolicy() (p PolicyTag, err error) {
	var webSite = 1
	policy, err := getLink(webSite, model.CategoryPrivatePolicy)
	if err != nil {
		return
	}
	tou, err := getLink(webSite, model.CategoryTermOfUse)
	if err != nil {
		return
	}
	p.PrivatePolicy = policy
	p.TermOfUse = tou
	return
}

func getLink(webSite int, category string) (link Link, err error) {
	t, err := model.GetResource(webSite, category)
	if err != nil {
		return
	}
	link = Link{
		Url:   t.Url,
		Name:  t.Name,
		Media: t.Media,
	}
	return
}

func getLinks(webSite int, category string) (links []Link, err error) {
	tmp, err := model.GetResources(webSite, category)
	for _, t := range tmp {
		menu := Link{
			Url:   t.Url,
			Name:  t.Name,
			Media: t.Media,
		}
		menu.Outer = t.OutLink == model.Link_outer
		links = append(links, menu)
	}
	return
}

func getSiteMap() (menus []Link, err error) {
	menus, err = getLinks(1, model.CategorySiteMap)
	return
}
