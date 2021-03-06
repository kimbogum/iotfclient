package main

import (
	"fmt"
	//import the Paho Go MQTT library
	"encoding/json"
	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
	"github.com/jimlawless/cfg"
	"log"
	"os"
	"time"
)

//define a function for the default message handler
var f MQTT.MessageHandler = func(client *MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func main() {
	//load the configuration
	configs := make(map[string]string)
	err := cfg.Load("iotfc.cfg", configs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", configs)

	//create a ClientOptions struct setting the broker address, clientid, turn
	//off trace output and set the default message handler
	opts := MQTT.NewClientOptions().AddBroker(configs["url"])
	opts.SetClientID("go-simple")
	opts.SetDefaultPublishHandler(f)

	//create and start a client using the above ClientOptions
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	//subscribe to the topic /go-mqtt/sample and request messages to be delivered
	//at a maximum qos of zero, wait for the receipt to confirm the subscription
	if token := c.Subscribe("go-mqtt/sample", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	//Publish 5 messages to /go-mqtt/sample at qos 1 and wait for the receipt
	//from the server after sending each message
	for i := 0; i < 5; i++ {
		text := fmt.Sprintf("this is msg #%d!", i)
		token := c.Publish("go-mqtt/sample", 0, false, text)
		token.Wait()
	}

	time.Sleep(3 * time.Second)

	//unsubscribe from /go-mqtt/sample
	if token := c.Unsubscribe("go-mqtt/sample"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	c.Disconnect(250)
}

func createMessage() []byte {
	message := D{MyName: "piem", Cputemp: 37.0, Cpuload: 2.0, Sine: 0.9}
	jsn, err := json.Marshal(message)
	if err != nil {
		log.Fatal(err)
	}
	return jsn
}

type Message struct {
	_id  string
	_rev string
	d    D
}

type D struct {
	MyName  string  `json:"myName"`
	Cputemp float32 `json:"cputemp"`
	Cpuload float32 `json:"cpuload"`
	Sine    float32 `json:"sine"`
}
