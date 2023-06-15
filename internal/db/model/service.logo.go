package model

import (
	"github.com/lishimeng/app-starter"
	persistence "github.com/lishimeng/go-orm"
)

// UpdateLogo 如果没有配置过logo, 会创建新logo的link
func UpdateLogo(webSiteId int, name, url, mediaPath string) (link Resource, err error) {

	err = app.GetOrm().Transaction(func(ctx persistence.TxContext) error {
		var tmp []Resource
		var e error
		_, e = ctx.Context.QueryTable(new(Resource)).
			Filter("Category", CategoryLogo).
			Filter("WebSiteId", webSiteId).
			All(&tmp)
		if e != nil {
			return e
		}
		if len(tmp) == 0 {
			link = Resource{
				WebSiteId: webSiteId,
				Name:      name,
				Url:       url,
				OutLink:   0,
				Category:  CategoryLogo,
				Index:     0,
			}
			_, e = ctx.Context.Insert(&link)
		} else {
			link = tmp[0]
			link.Name = name
			link.Url = url
			link.Media = mediaPath
			_, e = ctx.Context.Update(&link)
		}
		return e
	})
	return
}
