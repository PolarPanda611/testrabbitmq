/**
 * @ Author: Daniel Tan
 * @ Date: 2020-07-04 22:58:35
 * @ LastEditTime: 2020-07-05 00:25:56
 * @ LastEditors: Daniel Tan
 * @ Description:
 * @ FilePath: /testrabbitmq/publish.go
 * @
 */
package main

import "testrabbitmq/rabbitmq"

func main() {
	mq := rabbitmq.NewRabbitMqSimple("sample")
	mq.PublishSimple("test")
}
