package static

import (
	"embed"
)

//go:embed css/* fonts/* images/* js/* layout/* index.html
var Static embed.FS
