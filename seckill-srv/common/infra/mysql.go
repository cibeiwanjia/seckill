package infra

import (
	"fmt"
	"seckill/seckill-srv/common/config"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"seckill/seckill-srv/common/model"
)

func DBInit() {
	var err error
	data := config.AppCfg.Mysql
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		data.User, data.Password, data.Host, data.Port, data.Database)

	once := sync.Once{}
	once.Do(func() {
		config.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	})
	if err != nil {
		panic("mysql connect failed")
	}
	fmt.Println("mysql connect success")

	err = config.DB.AutoMigrate(
		&model.User{},
		&model.Product{},
		&model.SeckillOrder{},
		&model.SeckillProduct{},
	)
	if err != nil {
		panic("mysql migrate failed")
	}
	fmt.Println("mysql migrate success")

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := config.DB.DB()

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

}
