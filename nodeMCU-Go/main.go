package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	firmataAdaptor := firmata.NewAdaptor("10.10.100.107:3030")
	led := gpio.NewLedDriver(firmataAdaptor, "13")

	work := func() {
		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})
	}
	rebot := gobot.NewRobot("esp8266",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{led},
		work)
	rebot.Start()
}
