package mysql

import (
	"fmt"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

type PgsqlConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
}

func NewPgClient(config PgsqlConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", config.Host, config.User, config.Password, config.Database, config.Port)
	_db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return _db, err
}
