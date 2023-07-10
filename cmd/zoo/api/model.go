package api

type IndexPage struct {
	Layout
}

// Layout 页面结构
type Layout struct {
	Footer Footer
	Header Header
	Metas  Metas
	Banner Banner
}

type Metas struct {
	Keywords    string
	Description string
	Title       string
	Author      string
	Favicon     string
}

type Header struct {
	Menu Menu
	Logo Logo
}

type Logo struct {
	Link
}

type Banner struct {
	Title    string
	SubTitle string
	Image    string
	Links    []Link
}

// Footer 底部区域
type Footer struct {
	Copyright CopyrightTag
	Policy    PolicyTag

	Icp    Link
	Police Link

	CompanyDescribe CompanyDescribe
	CompanyAddress  CompanyAddress
	FriendlyLinks   FriendlyLinks
	SiteMap         SiteMap
	Links           []FooterLinks
}

type FooterLinks struct {
	Links Links
	Name  string
}

// CopyrightTag 版权
type CopyrightTag struct {
	Description string
	Link
}

// PolicyTag 免责
type PolicyTag struct {
	TermOfUse     Link
	PrivatePolicy Link
}

// Menu 菜单
type Menu struct {
	Items []Link
	Login Link
}

// MenuItem 菜单条目
type MenuItem struct {
	Link
}

type CompanyAddress struct {
	Addr  string
	Email string
	Call  string
}

type SiteMap struct {
	Links []Link
}

type FriendlyLinks struct {
	Links []Link
}

type Link struct {
	Url   string
	Name  string
	Media string
	Index int
	Outer bool // 外链
}

type CompanyDescribe struct {
	Logo
	Describe string
}
