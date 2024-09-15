package cache

import (
	"os"
)

func CreateCacheDir() {
	aletheiaCacheDir := GetAletheiaCache()
	os.Mkdir(aletheiaCacheDir, 0777)
}
