package env

import (
	"os"
)

func Get(key, fallback string) string {
	if os.Getenv(key) != "" {
		return os.Getenv(key)
	}

	return fallback
}
