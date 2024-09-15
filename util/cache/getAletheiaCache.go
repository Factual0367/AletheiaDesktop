package cache

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func GetAletheiaCache() string {
	userCachePath, err := os.UserCacheDir()
	if err != nil {
		log.Println(fmt.Sprintf("Error getting user cache: %s", err))
	}

	aletheiaCachePath := filepath.Join(userCachePath, "aletheia")
	return aletheiaCachePath
}
