package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseConfig struct {
	User       string
	Pass       string
	Database   string
	Protocol   string
	Connection string
}

type Resource struct {
	DB     *sql.DB
	config DatabaseConfig
}

func initializeConfig() DatabaseConfig {
	return DatabaseConfig{
		User:     os.Getenv("MYSQL_USER"),
		Pass:     os.Getenv("MYSQL_PASS"),
		Protocol: "tcp",
		Database: "sample",
		// Connection: "127.0.0.1:3306",
		Connection: os.Getenv("MYSQL_HOST"),
	}
}
func (d *DatabaseConfig) GetDataSourceName() string {
	log.Println(d.User, d.Pass, d.Protocol, d.Connection, d.Database)
	return fmt.Sprintf(
		"%s:%s@%s(%s)/%s?parseTime=true&loc=Asia%%2FTokyo&charset=utf8mb4",
		d.User, d.Pass, d.Protocol, d.Connection, d.Database,
	)
}

func NewResource() Resource {
	return Resource{
		config: initializeConfig(),
	}
}
func (resource *Resource) Initialize() {
	resource.config = initializeConfig()
	var err error
	resource.DB, err = NewDB("mysql", resource.config.GetDataSourceName())
	if err != nil {
		log.Println("接続失敗")
		panic(err)
	}
	log.Println("接続成功")
}

func (resource *Resource) Finalize() {
	if err := resource.DB.Close(); err != nil {
		panic(err)
	}
}

func NewDB(drivername, dataSourceName string) (*sql.DB, error) {
	return sql.Open(drivername, dataSourceName)
}
