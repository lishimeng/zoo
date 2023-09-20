package api

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/home/internal/db/model"
	"sort"
)

func getWebsite(name string) (ws model.EnterpriseWebSite, err error) {
	var tmp []model.EnterpriseWebSite
	_, err = app.GetOrm().Context.QueryTable(new(model.EnterpriseWebSite)).
		Filter("Name", name).
		All(&tmp)
	if len(tmp) > 0 {
		ws = tmp[0]
	}

	return
}

func getHeaderMenus(webSite int) (menus []Link, err error) {
	menus, err = getLinks(webSite, model.CategoryHeaderMenu)
	return
}

func getLogin(webSite int) (link Link, err error) {
	link, err = getLink(webSite, model.CategoryLogIn)
	return
}

func getLogo(webSite int) (link Link, err error) {
	link, err = getLink(webSite, model.CategoryLogo)
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
		menu.Outer = t.OutLink == model.LinkOuter
		links = append(links, menu)
	}
	return
}

type Links []Link

func (s Links) Len() int           { return len(s) }
func (s Links) Less(i, j int) bool { return s[i].Index < s[j].Index }
func (s Links) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func GetBannerLinks(webSite int) (links []Link, err error) {
	resources, err := model.GetResources(webSite, model.CategoryBanner)
	if err != nil {
		return
	}
	for _, r := range resources {
		var link = Link{
			Url:   r.Url,
			Name:  r.Name,
			Media: r.Media,
			Index: r.Index,
			Outer: r.OutLink == model.LinkOuter,
		}
		links = append(links, link)
	}
	return
}

func GetFooterLinks(webSite int) (footer []FooterLinks, err error) {
	resources, err := model.GetFooterResources(webSite)
	if err != nil {
		return
	}
	groups, err := model.GetFooterGroups(webSite)
	if err != nil {
		return
	}
	groupDetails, err := model.GetFooterGroupDetails(webSite)
	if err != nil {
		return
	}
	for _, g := range groups {
		var f = FooterLinks{Name: g.Name}
		for _, detail := range groupDetails {
			if detail.ResourceGroupId == g.Id {
				for _, r := range resources {
					if r.Id == detail.ResourceId {
						var link = Link{
							Url:   r.Url,
							Name:  r.Name,
							Media: r.Media,
							Index: r.Index,
							Outer: r.OutLink == model.LinkOuter,
						}
						f.Links = append(f.Links, link)
					}
				}
			}
		}
		sort.Sort(f.Links)
		footer = append(footer, f)
	}
	return
}
