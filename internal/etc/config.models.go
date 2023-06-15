package etc

type Configuration struct {
	Web web `toml:"web"`
	Db  db
}

type web struct {
	Listen string `toml:"listen"`
}

type db struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
	Ssl      string
}
