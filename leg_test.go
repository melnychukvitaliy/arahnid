package main

import "testing"

func TestEvaluateAngles(t *testing.T) {
	params := LegParams{6, 9, 12}
	testLeg := Leg{Angles{180, 180, 90}, params}
	cases := []struct {
		in   Position
		want Angles
	}{
		{Position{20, 2, 0}, Angles{95,28,180}},
	}
	for _, c := range cases {
		got := testLeg.EvaluateAngles(c.in)
		if got != c.want {
			t.Errorf("%+v.evaluateAngles == %+v\n, want %+v\n", c.in, got, c.want)
		}
	}
}
