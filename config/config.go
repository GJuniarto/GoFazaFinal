package config

import "fmt"

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func GetDBConfig() *DBConfig {
	return &DBConfig{
		Host:     "db.pinbduejflogwtdmqryb.supabase.co",
		Port:     5432,
		User:     "postgres",
		Password: "78p4cLXvgM0i56F0",
		DBName:   "postgres",
	}
}

func (config *DBConfig) GetDBURL() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DBName,
	)
}
