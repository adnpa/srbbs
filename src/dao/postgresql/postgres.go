package postgresql

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"srbbs/src/conf"
	"srbbs/src/query"
)

var db *gorm.DB

// Init 初始化Postgres连接
func init() {
	var err error
	cfg := conf.Cfg.PostgresConfig
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", cfg.Host, cfg.User, cfg.Password, cfg.DB, cfg.Port, cfg.SSLMode, cfg.TimeZone)
	db, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}
	db.AllowGlobalUpdate = true
	sqldb, _ := db.DB()
	sqldb.SetMaxOpenConns(cfg.MaxOpenConns)
	sqldb.SetMaxIdleConns(cfg.MaxIdleConns)
	// gorm设置db
	query.SetDefault(db)
	return
}

// Close 关闭数据库连接
func Close() {
	sqldb, _ := db.DB()
	sqldb.Close()
}

// 测试用，获取DB对象
func GetDB() *gorm.DB {
	return db
}
