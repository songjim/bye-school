package utiles

import "os"

func getEnv(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        value = defaultValue
    }

    return value
}