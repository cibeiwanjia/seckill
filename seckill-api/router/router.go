package router

import (
	"seckill/seckill-srv/common/pkg/middlewear"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.POST("/v1").Use(middlewear.AuthToken)
	return r
}
