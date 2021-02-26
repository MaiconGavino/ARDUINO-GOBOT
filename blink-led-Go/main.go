package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	firmataAdaptor := firmata.NewAdaptor("/COM3")
	led_Red := gpio.NewLedDriver(firmataAdaptor, "8")
	led_Blue := gpio.NewLedDriver(firmataAdaptor, "7")

	work := func() {
		gobot.Every(1*time.Second, func() {
			led_Blue.On()
			led_Red.Off()
			time.Sleep(500 * time.Millisecond)
			led_Blue.Off()
			led_Red.On()
			time.Sleep(500 * time.Millisecond)
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{led_Blue, led_Blue},
		work,
	)

	robot.Start()
}
