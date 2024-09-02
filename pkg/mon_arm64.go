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
	Pwd, _ = os.Getwd()
	NjMon  = filepath.Join(Pwd, "njmon")
)

func init() {
	if err := os.WriteFile(NjMon, njmon, 0755); err != nil {
		fmt.Println("write file error:", err)
		os.Exit(1)
	}
}
