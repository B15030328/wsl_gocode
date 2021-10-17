package def

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// 创建简单模式下rabbitmq实例
func NewRabbitMQSimple(queueName string) *Rabbitmq {
	rabbitmq := NewRabbitMQ(queueName, "", "")
	return rabbitmq
}

// 简单模式下生产代码
func (r *Rabbitmq) PublishSimple(message string) {

	// 声明队列
	_, err := r.Channel.QueueDeclare(
		r.QueueName,
		// 是否持久化
		false,
		// 是否自动删除
		false,
		// 是否具有排他性
		false,
		// 是否阻塞
		false,
		// 额外属性
		nil)
	if err != nil {
		fmt.Println(err)
	}

	r.Channel.Publish(r.Exchange, r.QueueName, false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

// 简单模式下消费代码
func (r *Rabbitmq) ConsumeSimple() {

	// 声明队列
	_, err := r.Channel.QueueDeclare(
		r.QueueName,
		// 是否持久化
		false,
		// 是否自动删除
		false,
		// 是否具有排他性
		false,
		// 是否阻塞
		false,
		// 额外属性
		nil)
	if err != nil {
		fmt.Println(err)
	}

	msgs, err := r.Channel.Consume(r.QueueName, "", true, false, false, false, nil)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)
	// 采用携程读消息
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
