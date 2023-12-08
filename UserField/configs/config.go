package Config

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	Email    EmailConfig    `json:"email"`
	Database DatabaseConfig `json:"database"`
}

type EmailConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type DatabaseConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DBName   string `json:"dbname"`
}

func ReadConfig() Config {
	file, err := os.Open("..config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
	return config
}
