package dao

import (
	"fmt"
	"github.com/otkinlife/go_db/driver"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var ConnPool *driver.DB

func init() {
	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	connectConfig := &driver.ConnectConfig{
		Host:     host,
		Port:     3306,
		DbName:   dbName,
		User:     user,
		Password: password,
		Charset:  "utf8mb4",
	}
	poolConfig := &driver.PoolConfig{
		MaxIdle:     300,
		MaxOpen:     500,
		MaxIdleTime: 2,
		MaxLifeTime: 30,
	}
	ConnPool = driver.NewDB(connectConfig, nil, poolConfig)
	err := ConnPool.RegisterDialectFunc(GetDialector).Pool()
	if err != nil {
		panic(err)
	}
}

func GetDialector(config *driver.ConnectConfig) gorm.Dialector {
	return mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.DbName, config.Charset))
}
