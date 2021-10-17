package main

import "rabbitmq/demo05_UseInProdEnv/def"

func main() {
	rb := def.NewRabbitMQSimple("chory")
	// rb.PublishSimple("hello consumer")
	// fmt.Println("发送成功")

	rb.ConsumeSimple()
}
