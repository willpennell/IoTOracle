package utils

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"sync"
)

var subWait = new(sync.WaitGroup)

var packedResult []byte

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	packedResult = msg.Payload()
	subWait.Done()
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

func StartMQTTClient(topic string) []byte {
	//defer w.Done()

	subWait.Add(1)
	opts := SetOpts()
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	sub(client, topic)
	client.Disconnect(250)
	return packedResult

	// TODO unpack result
	// TODO compare timestamp and make sure within sent timestamp
	// TODO add result IoTBool result struct
	// TODO convert to json string
	// TODO send back to smart contract
}

func sub(client mqtt.Client, topic string) {
	//defer wg.Done()
	token := client.Subscribe(topic, 0, messagePubHandler)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s\n", topic)
	subWait.Wait()
}
