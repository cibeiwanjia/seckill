package service

import (
	"context"
	"seckill/seckill-srv/basic/logic"
	"seckill/seckill-srv/protobuf"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	__.UnimplementedSeckillServerServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) SeckillCreate(ctx context.Context, in *__.PreSeckillCreateReq) (*__.PreSeckillCreateResp, error) {
	//抢单前置数据初始化
	logic.SeckillCreate(ctx, in)
	return &__.PreSeckillCreateResp{
		Success: true,
		Info:    "秒杀信息同步完毕",
	}, nil
}

func (s *Server) SearchStock(ctx context.Context, in *__.SearchStockReq) (*__.SearchStockResp, error) {
	logic.SearchStock(ctx, in)
	return &__.SearchStockResp{
		Success: true,
		Info:    "抢单成功",
	}, nil
}

func (s *Server) SeckillOrder(_ context.Context, in *__.SeckillOrderReq) (*__.SeckillOrderResp, error) {

	return &__.SeckillOrderResp{
		Success: true,
		Info:    "秒杀订单创建成功",
	}, nil
}
