package main

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var totalPubs int64 = 0
var totalSubs int64 = 0
var totalDuration int64 = 0

var connectStatu bool = true

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var messageSubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	totalSubs += 1
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {

	fmt.Printf("Connect lost: %v", err)
	connectStatu = false
}

func main() {
	var broker = "175.178.51.84"
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetAutoReconnect(true)
	opts.SetClientID("iot-mqtt-client-batch-insert-db")
	opts.SetUsername("lorawan_test")
	opts.SetPassword("Zy3K6PkSpGf43")
	opts.SetKeepAlive(60)
	opts.SetMaxReconnectInterval(5)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	client := mqtt.NewClient(opts)
	defer client.Disconnect(250)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	go func() {
		sub(client)
	}()

	for {
		if !connectStatu {
			println("程序退出")
			return
		}
		time.Sleep(time.Second * 5)
		fmt.Printf("当前消息队列入栈总数为%d\n", totalPubs)
		fmt.Printf("当前消息队列出栈总数为%d\n", totalPubs)
	}

}

func sub(client mqtt.Client) {
	topic := "xz/uplink"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s", topic)
}
