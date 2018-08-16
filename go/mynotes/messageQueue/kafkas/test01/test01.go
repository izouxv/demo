package main

import (
	"github.com/Shopify/sarama"
	"time"
	"log"
	"fmt"
	"os"
)

var Address = []string{"localhost:9092"}

func main()  {
	aaa()
	//syncProducer(Address...)
	//asyncProducer1(Address)
}

//同步消息模式
func syncProducer(address ...string)  {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5*time.Second
	p, err := sarama.NewSyncProducer(address, config)
	if err != nil {
		log.Printf("sarama.NewSyncProducer err, message=%s \n", err)
		return
	}
	defer p.Close()
	topic := "test"
	srcValue := "sync: this is a message. index=%d"
	for i:=0; i<10; i++ {
		value := fmt.Sprintf(srcValue, i)
		msg := &sarama.ProducerMessage{
			Topic:topic,
			Value:sarama.ByteEncoder(value),
		}
		part, offset, err := p.SendMessage(msg)
		if err != nil {
			log.Printf("send message(%s) err=%s \n", value, err)
		}else {
			fmt.Fprintf(os.Stdout, value + "发送成功，partition=%d, offset=%d \n", part, offset)
		}
		time.Sleep(2*time.Second)
	}
}

func aaa()  {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5*time.Second
	//config.Producer.RequiredAcks = sarama.WaitForAll
	//config.Producer.Partitioner = sarama.NewRandomPartitioner
	msg := &sarama.ProducerMessage{}
	msg.Topic = "test"
	//msg.Partition = int32(-1)
	//msg.Key = sarama.StringEncoder("key")
	msg.Value = sarama.ByteEncoder("hello.......")
	producer, err := sarama.NewSyncProducer(Address, config)
	if err != nil {
		fmt.Println("Failed to produce message:", err)
		os.Exit(500)
	}
	fmt.Println("pro:",producer)
	defer producer.Close()
	for i:=0;i>10;i++ {
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			fmt.Println("Failed to produce message: ", err)
		}
		fmt.Printf("partition=%d, offset=%d\n", partition, offset)
		time.Sleep(time.Second)
	}
}