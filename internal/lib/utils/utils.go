package utils

import (
	"os"
	"strconv"

	"github.com/spf13/viper"
)

func checkKey(key string) {
	_, envExists := os.LookupEnv(key)
	if !viper.IsSet(key) && !envExists {
		panic("Missing required key: " + key)
	}
}

func GetString(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return viper.GetString(key)
}

func GetStringOrPanic(key string) string {
	checkKey(key)
	return GetString(key)
}

func GetStringOrDefault(key string, def string) string {
	value := GetString(key)

	if value == "" {
		return def
	}

	return value
}

func GetIntOrDefault(key string, def int) int {
	value := GetString(key)

	if intValue, err := strconv.Atoi(value); err == nil {
		return intValue
	}

	return def
}

// func getIntOrPanic(key string) int {
// 	checkKey(key)
// 	return GetIntOrDefault(key, 0)
// }
