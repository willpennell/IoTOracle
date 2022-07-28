package utils

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"sync"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func SetOpts() *mqtt.ClientOptions {
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", GetBroker(), port))
	opts.SetClientID(GetClientID())
	opts.SetUsername(GetClientUsername())
	opts.SetPassword(GetClientPasswd())
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	//opts.SetProtocolVersion(4)
	return opts
}

func StartClient(topic string) {
	//defer w.Done()
	var wg sync.WaitGroup
	wg.Add(1)
	opts := SetOpts()
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	sub(client, topic, &wg)
	client.Disconnect(250)
	fmt.Println("here")
	wg.Wait()
}

func sub(client mqtt.Client, topic string, wg *sync.WaitGroup) {
	defer wg.Done()
	token := client.Subscribe(topic, 0, messagePubHandler)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s\n", topic)
	wg.Wait()
}
