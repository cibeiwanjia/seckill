package logic

import (
	"context"
	"seckill/seckill-srv/common/pkg"

	logger "github.com/cibeiwanjia/zapylx"
	"go.uber.org/zap"

	"seckill/seckill-srv/basic/dao"
	__ "seckill/seckill-srv/protobuf"
)

func SeckillCreate(ctx context.Context, in *__.PreSeckillCreateReq) {
	//判断商品是否在商品表中
	exist, err, goods := dao.ProductExist(in.ProductId)
	if err != nil {
		return
	}
	if !exist {
		logger.Info("该商品不存在",
			zap.Int("product_id", int(in.ProductId)))
	}
	//如果商品存在的话，将其加入秒杀配置表
	//就是加入mysql和redis
	err = dao.SeckillCreate(goods, in, ctx)
	if err != nil {
		return
	}
	//添加虚拟库存
	err = dao.SeckillRedisPush(ctx, in.ProductId, in.StockQuantity)
	if err != nil {
		return
	}
	logger.Info("进队列考虑流量削峰",
		zap.Int64("product_id", in.ProductId),
		zap.Int64("quantity", in.StockQuantity))
}

// 检查库存
func SearchStock(ctx context.Context, in *__.SearchStockReq) {
	//从redis中查找信息返回
	err := dao.SearchStock(ctx, in, in.ProductId)
	if err != nil {
		return
	}
	pkg.OrderGen("JD")
}
