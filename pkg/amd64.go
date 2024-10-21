//go:build amd64

package pkg

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
)

//go:embed amd64/njmon
var njmon []byte

var (
	pwd, _ = os.Getwd()
	Njmon  = ""
)

func init() {
	Njmon = filepath.Join(pwd, "njmon")
	if err := os.WriteFile(Njmon, njmon, 0755); err != nil {
		fmt.Println("Write file error:", err)
		os.Exit(1)
	}
}
