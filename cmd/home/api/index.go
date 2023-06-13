package api

type IndexPage struct {
	Footer Footer
	Header Header
}

type Header struct {
	Menu   Menu
	Logo   Logo
	Banner Banner
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

type Footer struct {
	Copyright CopyrightTag
	Policy    PolicyTag
}

type CopyrightTag struct {
	Description string
	Link        string
	Name        string
}

type PolicyTag struct {
	TermOfUse     string
	PrivatePolicy string
}

type Menu struct {
	Items []MenuItem
}

type MenuItem struct {
	Link   string
	Name   string
	Active bool // 激活(外链不可设置激活)
	Outer  bool // 外链
}
