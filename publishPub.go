/**
 * @ Author: Daniel Tan
 * @ Date: 2020-07-05 00:41:48
 * @ LastEditTime: 2020-07-05 00:42:00
 * @ LastEditors: Daniel Tan
 * @ Description:
 * @ FilePath: /testrabbitmq/publishPub.go
 * @
 */

package main

import "testrabbitmq/rabbitmq"

func main() {
	mq := rabbitmq.NewRabbitMqPub("sample")
	mq.PublishPub("test")
}
