package main

import (
        "gobot.io/x/gobot"
        "gobot.io/x/gobot/drivers/gpio"
        "gobot.io/x/gobot/platforms/raspi"
)

func main() {
        r := raspi.NewAdaptor()
        led := gpio.NewLedDriver(r, "7")

        work := func() {
                r.ServoWrite("7", 10)
        }

        robot := gobot.NewRobot("blinkBot",
                []gobot.Connection{r},
                []gobot.Device{led},
                work,
        )

        robot.Start()
}