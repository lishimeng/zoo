package model

import "github.com/lishimeng/app-starter"

type WebSite struct {
	app.Pk
	Name            string `orm:"column(name);null"`
	PageKeywords    string `orm:"column(page_keywords);null"`
	PageDescription string `orm:"column(page_description);null"`
	PageTitle       string `orm:"column(page_title);null"`
	PageAuthor      string `orm:"column(page_author);null"`
	BannerTitle     string `orm:"column(banner_title);null"`
	BannerSubTitle  string `orm:"column(banner_sub_title);null"`
	BannerMedia     string `orm:"column(banner_media);null"`
	Copyright       string `orm:"column(copyright);null"`
	CompanyDescribe string `orm:"column(company_describe);null"`
	CompanyAddress  string `orm:"column(company_address);null"`
	CompanyTel      string `orm:"column(company_tel);null"`
	CompanyEmail    string `orm:"column(company_email);null"`
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
}

type Menu struct {
}

const (
	CategoryHeaderMenu    = "header_menu"
	CategorySiteMap       = "site_map"
	CategoryFriendlyLink  = "friend_link"
	CategoryLogo          = "logo"
	CategoryLogIn         = "login"
	CategoryTermOfUse     = "term_of_use"
	CategoryPrivatePolicy = "private_policy"
	CategoryIcp           = "icp"
	CategoryPolice        = "police"
)

const (
	Link_Inner = iota
	Link_outer
)

func Tables() (t []interface{}) {
	t = append(t,
		new(WebSite),
		new(Resource),
	)
	return
}
