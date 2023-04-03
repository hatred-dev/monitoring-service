package configuration

import (
	"os"
	"strings"
)

func getEnvOrDefault(prefix string, key string, def string) string {

	prefixedKey := strings.ToUpper(prefix + key)
	var res string
	if res = os.Getenv(prefixedKey); res == "" {
		res = def
	}
	return res
}
