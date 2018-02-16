package main

import (
	"fmt"
	"math"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

// Leg is the basic structure for Bot&Co project
type Leg struct {
	initialHeight float64
	hipLength     float64
	shinLength    float64
}

func (leg Leg) EvaluateAngles(w, l, h float64) (hipAngle, shinAngle, rotationAngle float64) {
	s := leg.initialHeight
	b := leg.hipLength
	a := leg.shinLength

	ac := math.Hypot((s - h), math.Hypot(w, l))
	kc := math.Sqrt(w*w + l*l)

	shinAngle = math.Acos((b*b + a*a - ac*ac) / (2 * a * b))

	innerHipAngle := 0.0
	if s-h != 0 {
		math.Acos((math.Pow((s-h), 2) + ac*ac - kc*kc) / (2 * ac * (s - h)))
	}
	hipAngle = math.Acos((b*b+ac*ac-a*a)/(2*b*ac)) + innerHipAngle
	rotationAngle = math.Atan(l / w)

	return
}

func toDeg(a, b, c float64) (ad, bd, cd int) {
	ad = round(a * 180 / math.Pi)
	bd = round(b * 180 / math.Pi)
	cd = round(c * 180 / math.Pi)
	return
}

func round(f float64) int {
	return int(f + math.Copysign(0.5, f))
}

func main() {
	firmataAdaptor := firmata.NewAdaptor("/dev/ttyACM0")
	shinServo := gpio.NewServoDriver(firmataAdaptor, "9")
	hipServo := gpio.NewServoDriver(firmataAdaptor, "10")
	zServo := gpio.NewServoDriver(firmataAdaptor, "11")

	x, y, z := 5.0, 5.5, 5.0
	s, b, a := 4.7, 8.4, 12.5
	testLeg := Leg{s, b, a}

	fmt.Printf("Leg registered\n\n Initial height is %vcm;\nhip length = %vcm;\nshin length = %vcm\n\n", s, b, a)

	fmt.Printf("Trying to reach the point T (x = %v, y = %v, z = %v)\n", x, y, z)
	hip, shin, rotation := toDeg(testLeg.EvaluateAngles(x, y, z))
	fmt.Printf("Calculated angles: \n\tZ-rotation = %d\n\tHip angle = %d\n\tShin angle = %d\n", rotation, hip, shin)

	shin = 180
	// hip = 00
	// rotation = 0

	work := func() {
		shinServo.Move(uint8(180 - shin))
		hipServo.Move(uint8(hip))
		zServo.Move(uint8(rotation))

	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		work,
	)

	robot.Start()
}
