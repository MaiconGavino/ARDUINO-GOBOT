package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	firmataAdaptor := firmata.NewAdaptor("/COM3")
	led_Red := gpio.NewLedDriver(firmataAdaptor, "8")
	led_Blue := gpio.NewLedDriver(firmataAdaptor, "7")
	button := gpio.NewButtonDriver(firmataAdaptor, "2")

	work := func() {
		gobot.Every(1*time.Second, func() {

			button.On(gpio.ButtonPush, func(data interface{}) {
				fmt.Println("button pressed")
				led_Blue.Off()
				led_Red.On()
			})
			button.On(gpio.ButtonRelease, func(data interface{}) {
				fmt.Println("button release")
				led_Blue.On()
				led_Red.Off()
			})
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{led_Blue, led_Blue, button},
		work,
	)

	robot.Start()
}
