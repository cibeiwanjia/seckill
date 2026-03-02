package config

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	AppCfg *AppConf
	DB     *gorm.DB
	RDB    *redis.Client
)
