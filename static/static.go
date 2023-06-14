package static

import (
	"embed"
)

//go:embed assets/* vendor/* layout/* index.html
var Static embed.FS
