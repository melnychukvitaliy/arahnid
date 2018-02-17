package main

import (
	"fmt"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	firmataAdaptor := firmata.NewAdaptor("/dev/tty.usbmodem1431")
	skinServo := gpio.NewServoDriver(firmataAdaptor, "9")
	hipServo := gpio.NewServoDriver(firmataAdaptor, "10")

	work := func() {
		params := LegParams{6, 9, 12}
		testLeg := Leg{Angles{180, 180, 90}, params}
		angles := testLeg.EvaluateAngles(Position{10, 5, 0})
		fmt.Printf("I'm going to set up these angles %+v\n",angles )
		skinServo.Move(angles.skinAngle)
		hipServo.Move(angles.hipAngle)
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		work,
	)

	robot.Start()
}
