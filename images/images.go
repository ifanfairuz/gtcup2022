package images

import (
	"encoding/base64"
	"log"
	"os"
	"path"
)

func GetImageBgUri() string {
	bytes, err := os.ReadFile(path.Join("images", "bg.jpg"))
	if err != nil {
		log.Fatal(err)
		return ""
	}

	return "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(bytes)
}