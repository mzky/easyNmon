package pkg

import (
	"embed"
	"io/fs"
)

//go:embed assets
var Static embed.FS

func StaticFS() fs.FS {
	sfs, _ := fs.Sub(Static, "assets")
	return sfs
}
