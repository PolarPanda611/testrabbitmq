/**
 * @ Author: Daniel Tan
 * @ Date: 2020-07-05 00:24:58
 * @ LastEditTime: 2020-07-05 00:26:24
 * @ LastEditors: Daniel Tan
 * @ Description:
 * @ FilePath: /testrabbitmq/receive.go
 * @
 */
package main

import "testrabbitmq/rabbitmq"

func main() {
	mq := rabbitmq.NewRabbitMqSimple("sample")
	mq.ConsumeSimple()
}
