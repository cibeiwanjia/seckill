package infra

import (
	"seckill/seckill-srv/common/config"

	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func CacheInit() {
	data := config.AppCfg.Redis
	Addr := fmt.Sprintf("%s:%d", data.Host, data.Port)
	config.RDB = redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: data.Password, // no password set
		DB:       data.Database, // use default DB
	})
	err := config.RDB.Ping(context.Background()).Err()
	if err != nil {
		return
	}
	fmt.Println("redis连接成功")
}
