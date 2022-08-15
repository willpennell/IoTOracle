package utils

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"sync"
)

// allows WaitGroup for client so can stop once a result is received
var subWait = new(sync.WaitGroup)

var packedResult []byte

// messagePubHandler used for handling incoming messages from publisher (iot) device
var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	// message from topic
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	packedResult = msg.Payload() // add the message (json) to packedResult
	subWait.Done()               // will stop goroutine
}

// connectHandler used for connecting client
var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

// connectLostHandler used for lost connection
var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

// SetOpts creates and sets options for MQTT client
func SetOpts() *mqtt.ClientOptions {

	var port = 1883                                               // standard port for MQTT this could be in .env file
	opts := mqtt.NewClientOptions()                               // opts object
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", GetBroker(), port)) // add broker to connect to, info from .env
	opts.SetClientID(GetClientID())                               // set clientID, info from .env
	opts.SetUsername(GetClientUsername())                         // set client Username, info from .env
	opts.SetPassword(GetClientPasswd())                           // set client passwd, info from .env
	opts.SetDefaultPublishHandler(messagePubHandler)              // set publish handler. what happens when a message is recieved
	opts.OnConnect = connectHandler                               // what happens when connected
	opts.OnConnectionLost = connectLostHandler                    // what happens when connection lost
	//opts.SetProtocolVersion(4) // set options for MQTT protocol version use 3.1.1 as standard
	return opts
}

// StartMQTTClient connects to MQTT client and subscribes to particular topic
func StartMQTTClient(topic string) []byte {
	fmt.Println(topic)
	//defer w.Done()
	subWait.Add(1)                 // add a counter to global sync.WaitGroup function
	opts := SetOpts()              // create set options and return options object and assign to 'opts'
	client := mqtt.NewClient(opts) // start new MQTT client with opts
	// creates token when connected if token has error then throw panic
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	// subscription to topic
	sub(client, topic)     // pass in client and topic to subscribe
	client.Disconnect(250) // disconnects after 250 milliseconds but sync.Wait() in sub will pause until complete first
	return packedResult    // returns the message from pub
}

// sub function that subscribes to MQTT broker at requested topic
func sub(client mqtt.Client, topic string) {
	token := client.Subscribe(topic, 0, messagePubHandler) // publish handler callback used here
	fmt.Println("we are here")
	token.Wait()                                   // waits for a response from publish handler
	fmt.Printf("Subscribed to topic: %s\n", topic) // subscribed to topic
	subWait.Wait()                                 // waits for the counter to be decremented in publishHandler callback function
}
