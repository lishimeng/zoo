package api

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/home/internal/db/model"
	"sort"
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

func getPolice() (link Link, err error) {
	var webSite = 1
	link, err = getLink(webSite, model.CategoryPolice)
	return
}

func getIcp() (link Link, err error) {
	var webSite = 1
	link, err = getLink(webSite, model.CategoryIcp)
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

func GetFooterLinks() (footer []FooterLinks, err error) {
	var webSite = 1
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
							Outer: r.OutLink == 1,
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
