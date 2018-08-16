package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"os"
	"time"
)

//define a function for the default message handler
var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

var topic = "/test/test01"

func main() {
	//创建一个ClientOptions结构来设置代理地址, 客户端id, 关闭跟踪输出并设置默认消息处理程序
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetUsername("test01")
	opts.SetPassword("test01")
	opts.SetClientID("paho35521736665349")
	opts.SetDefaultPublishHandler(f)
	//使用上述客户端选项创建并启动客户端
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	//订阅主题并发送消息
	if token := c.Subscribe(topic, 1, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	//在qos 1发布5条消息到/go-mqtt/sample并等待收据,发送每条消息后从服务器发出
	//for i := 0; i < 5; i++ {
	//	text := fmt.Sprintf("this is msg #%d!", i)
	//	token := c.Publish(topic, 0, false, text)
	//	token.Wait()
	//}
	time.Sleep(30 * time.Second)
	//取消订阅主题
	if token := c.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	c.Disconnect(250)
}