package SimlpePublish

import (
	"fmt"

	"seckill/seckill-srv/common/rabbitmq/RabbitMQ"
)

func SimplePublish(msg string) {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" +
		"yh")
	rabbitmq.PublishSimple(msg)
	fmt.Println("发送成功！")
}
