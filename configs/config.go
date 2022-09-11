package configs

import "os"

var config = map[string]string{
	"APP_NAME":  "Golearn",
	"HTTP_PORT": "8080",

	//db configs
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
