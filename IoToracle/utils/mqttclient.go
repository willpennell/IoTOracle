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
	MQTTMESSAGE(fmt.Sprintf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic()))
	packedResult = msg.Payload() // add the message (json) to packedResult
	subWait.Done()               // will stop goroutine
}

// connectHandler used for connecting client
var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	MQTTMESSAGE("connected..")
}

// connectLostHandler used for lost connection
var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	MQTTMESSAGE(fmt.Sprintf("Connect lost: %v", err))
}

// SetOpts creates and sets options for MQTT client
func SetOpts(clientID string) *mqtt.ClientOptions {

	var port = 1883                                               // standard port for MQTT this could be in .env file
	opts := mqtt.NewClientOptions()                               // opts object
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", GetBroker(), port)) // add broker to connect to, info from .env
	opts.SetClientID(clientID)                                    // set clientID, info from .env
	opts.SetUsername(GetClientUsername())                         // set client Username, info from .env
	opts.SetPassword(GetClientPasswd())                           // set client passwd, info from .env
	opts.SetDefaultPublishHandler(messagePubHandler)              // set publish handler. what happens when a message is recieved
	opts.OnConnect = connectHandler                               // what happens when connected
	opts.OnConnectionLost = connectLostHandler                    // what happens when connection lost
	//opts.SetProtocolVersion(4) // set options for MQTT protocol version use 3.1.1 as standard
	return opts
}

var client1 mqtt.Client

func setClient(client mqtt.Client) {
	client1 = client
}

// StartMQTTClient connects to MQTT client and subscribes to particular topic
func StartMQTTClient(id uint64, topic string, clientID string) []byte {
	fmt.Println(topic)
	//defer w.Done()
	subWait.Add(1)                 // add a counter to global sync.WaitGroup function
	opts := SetOpts(clientID)      // create set options and return options object and assign to 'opts'
	client := mqtt.NewClient(opts) // start new MQTT client with opts
	// creates token when connected if token has error then throw panic
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	// subscription to topic
	setClient(client)
	sub(client, topic) // pass in client and topic to subscribe
	//client.Unsubscribe(topic)
	var tenSecs uint = 10 * 1000
	client.Disconnect(uint((Requests[id].ElapsedTime*2)*1000) + tenSecs) // disconnects after 250 milliseconds but sync.Wait() in sub will pause until complete first
	return packedResult                                                  // returns the message from pub
}

// sub function that subscribes to MQTT broker at requested topic
func sub(client mqtt.Client, topic string) {
	token := client.Subscribe(topic, 0, messagePubHandler) // publish handler callback used here

	token.Wait()         // waits for a response from publish handler
	MQTTBROKERSUB(topic) // subscribed to topic
	subWait.Wait()       // waits for the counter to be decremented in publishHandler callback function
}

func forceCloseConnection() {
	fmt.Println("Closing connection, no longer required.")
	client1.Disconnect(250)
}
