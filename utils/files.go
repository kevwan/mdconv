package utils

import (
	"fmt"
	"path/filepath"
	"strings"
)

func ChangeExt(filename, ext string) string {
	name := strings.TrimSuffix(filename, filepath.Ext(filename))
	return fmt.Sprintf("%s.%s", name, ext)
}
