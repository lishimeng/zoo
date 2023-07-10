package model

import (
	"github.com/lishimeng/app-starter"
	persistence "github.com/lishimeng/go-orm"
)

func GetResource(webSite int, category string) (link Resource, err error) {
	var tmp []Resource
	_, err = app.GetOrm().Context.QueryTable(new(Resource)).
		Filter("WebSiteId", webSite).
		Filter("Category", category).
		All(&tmp)
	if len(tmp) > 0 {
		link = tmp[0]
	}

	return
}

func GetResources(webSite int, category string) (link []Resource, err error) {
	_, err = app.GetOrm().Context.QueryTable(new(Resource)).
		Filter("WebSiteId", webSite).
		Filter("Category", category).
		OrderBy("index").
		All(&link)
	return
}

func AddLink(ctx persistence.TxContext, d Resource, isOuterLink bool) (link Resource, err error) {

	if isOuterLink {
		link.OutLink = LinkOuter
	} else {
		link.OutLink = LinkInner
	}
	link.WebSiteId = d.WebSiteId
	link.Name = d.Name
	link.Url = d.Url
	link.Media = d.Media
	link.Category = d.Category
	link.Index = d.Index

	_, err = ctx.Context.Insert(&link)
	return
}

func UpdateLink(ctx persistence.TxContext, d Resource) (link Resource, err error) {

	link.Pk.Id = d.Id
	link.Name = d.Name
	link.Url = d.Url
	link.OutLink = d.OutLink
	link.Media = d.Media
	link.Index = d.Index

	_, err = ctx.Context.Update(&link, "Name", "Url", "Outer", "Media", "Index")
	return
}

func DeleteLink(ctx persistence.TxContext, id int) (link Resource, err error) {
	link.Pk.Id = id
	_, err = ctx.Context.Delete(&link, "id")
	return
}
