/**
 * @ Author: Daniel Tan
 * @ Date: 2020-07-04 23:05:37
 * @ LastEditTime: 2020-07-05 00:49:50
 * @ LastEditors: Daniel Tan
 * @ Description:
 * @ FilePath: /testrabbitmq/rabbitmq/rabbitmq.go
 * @ Daniel Tan Modified
 */

package rabbitmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

var MqURL string = "amqp://testuser:testuser@127.0.0.1:5672/test"

// RabbitMq instance for rabbitmq
type RabbitMq struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	QueueName string
	Exchange  string
	Key       string
	MqURL     string
}

// NewRabbitMqSimple new mq simple
func NewRabbitMqSimple(queueName string) *RabbitMq {
	return NewRabbitMq(queueName, "", "", MqURL)
}

// NewRabbitMqPub new mq simple
func NewRabbitMqPub(exchangeName string) *RabbitMq {
	return NewRabbitMq("", exchangeName, "", MqURL)
}

// NewRabbitMq factory for rabbitmq
func NewRabbitMq(queueName, exchange, key, mqURL string) *RabbitMq {
	conn, err := amqp.Dial(mqURL)
	if err != nil {
		log.Fatal(err)
	}
	channel, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	return &RabbitMq{
		conn:      conn,
		channel:   channel,
		QueueName: queueName,
		Exchange:  exchange,
		Key:       key,
		MqURL:     mqURL,
	}
}

// Destroy instance
/**
 * @ Author: Daniel Tan
 * @ Date: 2020-07-05 00:04:38
 * @ LastEditTime: Do not edit
 * @ LastEditors: Daniel Tan
 * @ description: Destroy instance for rabbitmq
 */
func (r *RabbitMq) Destroy() {
	r.channel.Close()
	r.conn.Close()
}

func (r *RabbitMq) initQueue() error {
	_, err := r.channel.QueueDeclare(
		r.QueueName, // 队列名称
		false,       // 是否持久化
		false,       // 是否自动删除
		false,       //是否排他
		false,       // 是否阻塞
		nil,         //是否其他属性
	)
	return err
}

func (r *RabbitMq) initExchange() error {
	// 订阅模式
	return r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout", // fanout 广播模式
		true,
		false,
		false,
		false,
		nil,
	)
}

// PublishSimple publish simple
/**
 * @ Author: Daniel Tan
 * @ Date: 2020-07-04 23:41:25
 * @ LastEditTime: Do not edit
 * @ LastEditors: Daniel Tan
 * @ description:
 * @ param message , the message you want to publish
 * @ return:
 */
func (r *RabbitMq) PublishSimple(messsage string) {
	// 1. 注册队列
	err := r.initQueue()
	if err != nil {
		log.Fatal(err)
	}
	// 发送消息
	r.channel.Publish(
		r.Exchange,  // 交换机名称
		r.QueueName, //队列名称
		false,       // if true ,如果发送失败则会回退
		false,       // if true , 没有消费者则消息回退
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(messsage),
		},
	)
}

func (r *RabbitMq) PublishPub(messsage string) {
	// 1. 注册队列
	err := r.initExchange()
	if err != nil {
		log.Fatal(err)
	}
	// 发送消息
	r.channel.Publish(
		r.Exchange, // 交换机名称
		"",
		false, // if true ,如果发送失败则会回退
		false, // if true , 没有消费者则消息回退
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(messsage),
		},
	)
}

// ConsumeSimple consume simple message
/**
 * @ Author: Daniel Tan
 * @ Date: 2020-07-05 00:04:22
 * @ LastEditTime: Do not edit
 * @ LastEditors: Daniel Tan
 * @ description:
 */
func (r *RabbitMq) ConsumeSimple() {
	err := r.initQueue()
	if err != nil {
		log.Fatal(err)
	}
	msgs, err := r.channel.Consume(
		r.QueueName,
		"",    // 区分多个消费者
		true,  //是否自动应答
		false, //是否排他
		false, // if true ， 消息无法传递给同一个connection中的消费者
		false, // 是否阻塞
		nil,   //其他属性
	)
	if err != nil {
		log.Fatal(err)
	}
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			log.Printf("receive msg : %v", string(msg.Body))
		}
	}()
	fmt.Println("listening msg from queue")
	<-forever
}

func (r *RabbitMq) ConsumePub() {
	//初始化交换机
	err := r.initExchange()
	if err != nil {
		log.Fatal(err)
	}
	//
	q, err := r.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	err = r.channel.QueueBind(q.Name, "", r.Exchange, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	msgs, err := r.channel.Consume(
		q.Name,
		"",    // 区分多个消费者
		true,  //是否自动应答
		false, //是否排他
		false, // if true ， 消息无法传递给同一个connection中的消费者
		false, // 是否阻塞
		nil,   //其他属性
	)
	if err != nil {
		log.Fatal(err)
	}
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			log.Printf("receive msg : %v", string(msg.Body))
		}
	}()
	fmt.Println("listening msg from queue")
	<-forever
}
