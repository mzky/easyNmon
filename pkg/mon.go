package pkg

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
)

//go:embed njmon
var njmon []byte

var (
	Pwd, _ = os.Getwd()
	NjMon  = filepath.Join(Pwd, "njmon")
)

func init() {
	if err := os.WriteFile(NjMon, njmon, 0755); err != nil {
		fmt.Println("write file error:", err)
	}
}
