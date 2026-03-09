package config

import "os"

func GetEnv(key, defaultVal string) string {
	v := os.Getenv(key)
	if v == "" {
		return defaultVal
	}

	return v
}
