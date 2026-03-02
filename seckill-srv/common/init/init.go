package init

import "seckill/seckill-srv/common/infra"

func init() {
	ConfigInit()
	infra.DBInit()
	infra.CacheInit()
}
