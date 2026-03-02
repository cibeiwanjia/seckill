package dao

import (
	"context"
	"errors"
	"fmt"
	"seckill/seckill-srv/common/config"
	"seckill/seckill-srv/common/model"
	"seckill/seckill-srv/common/pkg/middlewear"

	"seckill/seckill-srv/protobuf"
)

// 查看商品是否存在于商品表中
func ProductExist(productId int64) (bool, error, model.Product) {
	var product model.Product
	if err := config.DB.Debug().Where("id = ?", productId).Limit(1).Find(&product).Error; err != nil {
		return false, errors.New("数据库查询失败"), model.Product{}
	}
	if product.ID == 0 {
		return false, errors.New("该商品不存在"), model.Product{}
	}
	return true, nil, product
}

// 将商品加入秒杀配置表
func SeckillCreate(product model.Product, in *__.PreSeckillCreateReq, ctx context.Context) error {
	//先将其加入mysql，然后将其加入redis
	seckillProduct := &model.SeckillProduct{
		ProductID:    int(in.ProductId),
		SeckillPrice: product.Price,
		SeckillNum:   int(in.StockQuantity),
		StratTime:    int(in.StartTime),
		EndTime:      int(in.EndTime),
		MaxPerLimit:  int(in.MaxPerLimit),
	}
	if err := config.DB.Debug().Create(&seckillProduct).Error; err != nil {
		return errors.New("加入秒杀商品失败")
	}
	fmt.Println("加入秒杀商品成功")

	seckillKey := fmt.Sprintf("seckill:%d", in.ProductId)
	err := config.RDB.HMSet(ctx, seckillKey, seckillProduct).Err()
	if err != nil {
		return errors.New("秒杀商品入reids失败")
	}
	return nil
}

// 添加虚拟库存
func SeckillRedisPush(ctx context.Context, productID int64, quantity int64) error {
	listKey := fmt.Sprintf("listKey:%d", productID)
	for i := 0; i < int(quantity); i++ {
		err := config.RDB.LPush(ctx, listKey, productID).Err()
		if err != nil {
			return errors.New("虚拟库存添加失败")
		}
	}
	middlewear.GenerateAccessToken(productID)
	return nil
}

func SearchStock(ctx context.Context, in *__.SearchStockReq, productID int64) error {
	listKey := fmt.Sprintf("listKey:%d", productID)
	err := config.RDB.Exists(ctx, listKey).Err()
	if err != nil {
		return errors.New("抢单成功")
	}
	return nil
}
