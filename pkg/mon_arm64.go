//go:build arm64

package pkg

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
)

//go:embed arm64/njmon
var njmon []byte

var (
	pwd, _ = os.Getwd()
	Njmon  = filepath.Join(pwd, "njmon")
)

func init() {
	if err := os.WriteFile(Njmon, njmon, 0755); err != nil {
		fmt.Println("Write file error:", err)
		os.Exit(1)
	}
}
