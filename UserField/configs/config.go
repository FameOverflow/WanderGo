package Config

import (
	"os"
	"strconv"
)

type Config struct {
	Email    EmailConfig    `json:"email"`
	Database DBConfig `json:"database"`
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

func getIntEnv(key string, defaultValue int) int {
	// 从环境变量中获取整数值，如果未设置则使用默认值
	if value, err := strconv.Atoi(os.Getenv(key)); err == nil {
		return value
	}
	return defaultValue
}

func ReadDBConfig() DBConfig {
	// 从环境变量中读取数据库配置信息
	return DBConfig{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     getIntEnv("DB_PORT", 3306),
		DBName:   os.Getenv("DB_NAME"),
	}
}

func ReadEMConfig() EmailConfig {
	// 从环境变量中读取数据库配置信息
	return EmailConfig{
		Host:     os.Getenv("EMAIL_HOST"),
		Port:     getIntEnv("EMAIL_PORT", 587),
		UserName: os.Getenv("EMAIL_USERNAME"),
		Password: os.Getenv("EMAIL_PASSWORD"),
	}
}

var EMConfig EmailConfig = ReadEMConfig()
var DBconfig DBConfig = ReadDBConfig()