package filemanipulation

import (
	"io/fs"
	"os"
)

func WriteFile(file string) {
	if err := os.Mkdir("images", fs.FileMode(os.O_WRONLY)); err == nil {
		os.CreateTemp("image/", "image*.web")
	}
}
