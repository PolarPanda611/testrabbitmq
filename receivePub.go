/**
 * @ Author: Daniel Tan
 * @ Date: 2020-07-05 00:44:28
 * @ LastEditTime: 2020-07-05 00:49:59
 * @ LastEditors: Daniel Tan
 * @ Description:
 * @ FilePath: /testrabbitmq/receivePub.go
 * @
 */
package main

import "testrabbitmq/rabbitmq"

func main() {
	mq := rabbitmq.NewRabbitMqPub("sample")
	mq.ConsumePub()
}
