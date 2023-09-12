package model

import "github.com/lishimeng/app-starter"
import "github.com/lishimeng/app-starter/cms"

type EnterpriseWebSite struct {
	app.Pk
	Name            string `orm:"column(name);null"`
	PageKeywords    string `orm:"column(page_keywords);null"`
	PageDescription string `orm:"column(page_description);null"`
	PageTitle       string `orm:"column(page_title);null"`
	PageAuthor      string `orm:"column(page_author);null"`
	BannerTitle     string `orm:"column(banner_title);null"`
	BannerSubTitle  string `orm:"column(banner_sub_title);null"`
	BannerMedia     string `orm:"column(banner_media);null"`
	CompanyDescribe string `orm:"column(company_describe);null"`
	CompanyAddress  string `orm:"column(company_address);null"`
	CompanyTel      string `orm:"column(company_tel);null"`
	CompanyEmail    string `orm:"column(company_email);null"`
	app.TableChangeInfo
}

type Resource struct {
	app.Pk
	WebSiteId int    `orm:"column(web_site);null"`
	Name      string `orm:"column(name);null"`
	Url       string `orm:"column(url);null"`
	OutLink   int    `orm:"column(out_link);null"`
	Media     string `orm:"column(media);null"`
	Category  string `orm:"column(category);null"` // 类型
	Index     int    `orm:"column(index);null"`    // 排序
	app.TableChangeInfo
}

type ResourceGroup struct {
	app.Pk
	WebSiteId int    `orm:"column(web_site);null"`
	Name      string `orm:"column(name);null"`
	Index     int    `orm:"column(index);null"` // 排序
	app.TableChangeInfo
}

type GroupDetail struct {
	app.Pk
	WebSiteId       int `orm:"column(web_site);null"`
	ResourceId      int `orm:"column(resource_id);null"`
	ResourceGroupId int `orm:"column(resource_group_id);null"`
	app.TableInfo
}

const (
	CategoryHeaderMenu    = "header_menu"
	CategoryLogo          = "logo"
	CategoryLogIn         = "login"
	CategoryBanner        = "banner"
	CategoryTermOfUse     = "term_of_use"
	CategoryPrivatePolicy = "private_policy"
	CategoryIcp           = "icp"
	CategoryPolice        = "police"
	CategoryFooter        = "footer"
)

const (
	LinkInner = iota
	LinkOuter
)

func Tables() (t []interface{}) {
	t = append(t,
		new(EnterpriseWebSite),
		new(Resource),
		new(ResourceGroup),
		new(GroupDetail),
		new(cms.WebSite),
	)
	return
}
