package main

import (
	"math"
)

// Leg is the basic structure for Bot&Co project
type Leg struct {
	angles Angles
	params LegParams
}

//LegParams is the basic leg configuration that measures in santimetres
type LegParams struct {
	heigth,
	hipLength,
	skinLength float64
}

// Position contains all coordinates that measures in santimetres
type Position struct {
	x, y, z float64
}

// Angles is base structure that store all angles to controll leg
type Angles struct {
	hipAngle,
	skinAngle,
	rotationAngle uint8
}

func round(f float64) int {
	return int(f + math.Copysign(0.5, f))
}

func toDeg(a float64) (ad uint8) {
	return uint8(round(a * 180 / math.Pi))

}

// EvaluateAngles calculate angles that servo should scroll
func (leg Leg) EvaluateAngles(position Position) (angles Angles) {
	pointHeigthDiff := math.Abs(leg.params.heigth - position.y)
	a := leg.params.hipLength
	b := math.Hypot(pointHeigthDiff, position.x)
	c := leg.params.skinLength

	cosA := (a*a + c*c - b*b) / (2 * a * c)
	cosB := (a*a + b*b - c*c) / (2 * a * b)
	skinAngle := 180 - toDeg(math.Acos(cosA))
	var hipAngle uint8;
	if leg.params.heigth < position.y {
		cosP := (position.x / b)
		pAngle := toDeg(math.Acos(cosP))
		hipAngle = toDeg(math.Acos(cosB)) + pAngle + 90
	} else {
		cosP := (pointHeigthDiff / b)
		pAngle := toDeg(math.Acos(cosP))
		hipAngle = toDeg(math.Acos(cosB)) + pAngle
	}
	return Angles{hipAngle, skinAngle, 180}
}
