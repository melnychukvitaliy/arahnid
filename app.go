package main

import (
        "fmt"
        "time"
        "gobot.io/x/gobot"
        "gobot.io/x/gobot/drivers/gpio"
        "gobot.io/x/gobot/platforms/raspi"
)

func main() {
        raspiAdaptor := raspi.NewAdaptor()
        servo := gpio.NewServoDriver(raspiAdaptor, "7")

        work := func() {
                gobot.Every(1*time.Second, func() {
                        i := uint8(gobot.Rand(180))
                        fmt.Println("Turning", i)
                        servo.Move(i)
                })
        }

        robot := gobot.NewRobot("servoBot",
                []gobot.Connection{raspiAdaptor},
                []gobot.Device{servo},
                work,
        )
        robot.Start()
}