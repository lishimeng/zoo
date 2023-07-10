package etc

type Configuration struct {
	Web web `toml:"web"`
	Db  db  `toml:"db"`
}

type web struct {
	Listen    string `toml:"listen"`
	Cache     int    `toml:"cache"`
	WebSiteId int    `toml:"webSiteId"`
}

type db struct {
	User     string `toml:"user"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Database string `toml:"database"`
	Ssl      string `toml:"ssl"`
}
