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
}

type Header struct {
	Menu Menu
	Logo Logo
}

type Logo struct {
	Title string
	Url   string
	Image string
}

type Banner struct {
	Title    string
	SubTitle string
	Image    string
}

// Footer 底部区域
type Footer struct {
	Copyright CopyrightTag
	Policy    PolicyTag
}

// CopyrightTag 版权
type CopyrightTag struct {
	Description string
	Link        string
	Name        string
}

// PolicyTag 免责
type PolicyTag struct {
	TermOfUse     string
	PrivatePolicy string
}

// Menu 菜单
type Menu struct {
	Items []MenuItem
	Login string
}

// MenuItem 菜单条目
type MenuItem struct {
	Link   string
	Name   string
	Active bool // 激活(外链不可设置激活)
	Outer  bool // 外链
}
