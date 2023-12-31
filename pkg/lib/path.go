package lib

import (
	"os/user"
	"path/filepath"
	"strings"
)

func MustPath(path string) string {
	if u, err := user.Current(); err == nil {
		if path == "~" {
			return u.HomeDir
		} else if strings.HasPrefix(path, "~/") {
			return filepath.Join(u.HomeDir, path[2:])
		}
	}
	return path
}
