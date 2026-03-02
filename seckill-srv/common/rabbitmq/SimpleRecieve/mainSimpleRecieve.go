package SimpleRecieve

import (
	"log"
	"seckill/seckill-srv/common/rabbitmq/RabbitMQ"
)

func SimpleReceive() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" +
		"yh")
	msgs, err := rabbitmq.ConsumeSimple()
	if err != nil {
		return
	}

	forever := make(chan bool)
	//启用协程处理消息
	go func() {
		for d := range msgs {
			//消息逻辑处理，可以自行设计逻辑
			log.Printf("Received a message: %s", d.Body)

		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
