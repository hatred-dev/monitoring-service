package configuration

import (
    "fmt"
    "os"
    "strings"
)

func str(s string) *string {
    return &s
}
func getEnvOrDefault(prefix string, key string, def *string) string {
    prefixedKey := strings.ToUpper(prefix + key)
    res := os.Getenv(prefixedKey)
    if res == "" {
        if def == nil {
            panic(fmt.Sprintf("environment variable %s not set", prefixedKey))
        }
        res = *def
    }
    return res
}
