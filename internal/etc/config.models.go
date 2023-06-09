package etc

type Configuration struct {
	Web web `toml:"web"`
}

type web struct {
	Listen string `toml:"listen"`
}
