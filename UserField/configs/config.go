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

type DBConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DBName   string `json:"dbname"`
}

