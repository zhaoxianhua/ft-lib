package tdengine

import (
	"database/sql"
	"fmt"

	_ "github.com/taosdata/driver-go/v3/taosSql"
)

type TdEngineConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
}

func NewTdClient(td TdEngineConfig) (*sql.DB, error) {
	var taosDsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", td.User, td.Password, td.Host, td.Port, td.Database)
	taos, err := sql.Open("taosSql", taosDsn)
	if err != nil {
		fmt.Println("failed to connect TDengine, err:", err)
		return nil, err
	}

	err = taos.Ping()
	return taos, err
}
