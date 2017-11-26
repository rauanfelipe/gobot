package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/mqtt"
)

func main() {
	// Public mqtt broker
	mqttAdaptor := mqtt.NewAdaptor("tcp://test.mosquitto.org:1883", "")

	work := func() {
		// Subscribe
		mqttAdaptor.On("my_topic", func(msg mqtt.Message) {
			s := string(msg.Payload())
			fmt.Printf("%s: %s\n", msg.Topic(), s)
		})

		// Publish
		count := 0
		gobot.Every(1*time.Second, func() {
			msg := fmt.Sprintf("%s_%d", "message", count)
			data := []byte(msg)

			mqttAdaptor.Publish("my_topic", data)

			count++
		})
	}

	robot := gobot.NewRobot("mqttBot",
		[]gobot.Connection{mqttAdaptor},
		work)

	robot.Start()
}
