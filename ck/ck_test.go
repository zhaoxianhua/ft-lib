package ck

import (
	"fmt"
	gormck "gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	"log"
	"testing"
)

func TestCK(t *testing.T) {
	ckConfig := ClickHouseConfig{
		Host:     "172.16.2.143",
		Port:     9000,
		User:     "default",
		Password: "123456",
		Database: "etl",
	}
	conn := NewCkClient(ckConfig)
	db, err := gorm.Open(gormck.New(gormck.Config{
		Conn: conn, // initialize with existing database conn
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m := make([]map[string]interface{}, 0)
	err = db.Table("log_alarm").Where("log_time between ? and ?", 1690992000, 1693670400).Where("log_type = ?", 3).Offset(10).Limit(100).Find(&m).Error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)
}
