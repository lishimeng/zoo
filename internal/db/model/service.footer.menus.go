package model

import "github.com/lishimeng/app-starter"

func GetFooterGroups(webSite int) (gp []ResourceGroup, err error) {
	_, err = app.GetOrm().Context.
		QueryTable(new(ResourceGroup)).
		Filter("WebSiteId", webSite).
		OrderBy("Index").All(&gp)
	return
}

func GetFooterResources(webSite int) (r []Resource, err error) {
	r, err = GetResources(webSite, CategoryFooter)
	return
}

func GetFooterGroupDetails(webSite int) (gd []GroupDetail, err error) {
	_, err = app.GetOrm().Context.
		QueryTable(new(GroupDetail)).
		Filter("WebSiteId", webSite).All(&gd)
	return
}
