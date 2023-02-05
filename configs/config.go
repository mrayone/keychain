package configs

import (
	"os"
	"strconv"
	"time"
)

var config = map[string]string{
	"APP_NAME": "keychain",

	// http server
	"HTTP_PORT":                        "9000",
	"HTTP_SERVER_READ_TIMEOUT_MILLIS":  "60000",
	"HTTP_SERVER_WRITE_TIMEOUT_MILLIS": "60000",

	// db configs
	"APP_DB_USERNAME": "postgres",
	"APP_DB_PASSWORD": "postgres",
	"APP_DB_NAME":     "sales",
}

func GetString(k string) string {
	v := os.Getenv(k)
	if v == "" {
		return config[k]
	}

	return v
}

// GetInt value of a given env var
func GetInt(k string) int {
	v := GetString(k)
	i, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}

	return i
}

// GetDuration value of a given env var
func GetDuration(k string) time.Duration {
	return time.Duration(GetInt(k)) * time.Millisecond
}
