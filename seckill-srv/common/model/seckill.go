package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null;comment:用户名"`
	Password string `gorm:"type:varchar(50);not null;comment:密码"`
	Phone    string `gorm:"type:varchar(20);not null;comment:电话号码"`
}
type Product struct {
	gorm.Model
	GoodsName string `gorm:"type:varchar(100);not null;comment:商品名"`
	Stock     int    `gorm:"type:int;not null;comment:库存"`
	Price     string `gorm:"type:varchar(10);not null;comment:价格"`
	ImgUrl    string `gorm:"type:varchar(500);not null;comment:商品主图"`
}
type SeckillProduct struct {
	gorm.Model
	ProductID    int    `gorm:"type:int;not null;comment:商品id"`
	SeckillPrice string `gorm:"type:varchar(10);not null;comment:秒杀价"`
	SeckillNum   int    `gorm:"type:int;not null;comment:秒杀商品数量"`
	StratTime    int    `gorm:"type:int;not null;comment:秒杀开始时间"`
	EndTime      int    `gorm:"type:int;not null;comment:秒杀结束时间"`
	MaxPerLimit  int    `gorm:"type:int;not null;comment:限购"`
}
type SeckillOrder struct {
	gorm.Model
	OrderSn    string `gorm:"type:varchar(30);not null;comment:订单编号"`
	UserID     int    `gorm:"type:int;not null;comment:用户ID"`
	ProductID  int    `gorm:"type:int;not null;comment:商品id"`
	Quantity   int    `gorm:"type:int;not null;comment:购买数量"`
	TotalPrice string `gorm:"type:varchar(12);not null;comment:订单总金额"`
}
