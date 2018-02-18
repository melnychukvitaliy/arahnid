package main

import "testing"

func TestEvaluateAngles(t *testing.T) {
	params := LegParams{10, 9, 12}
	testLeg := Leg{Angles{180, 180, 90}, params}
	cases := []struct {
		in   Position
		want Angles
	}{
		{Position{10, 5, 0}, Angles{135,118,180}},
		{Position{10, 12, 0}, Angles{178,124,180}},
	}
	for _, c := range cases {
		got := testLeg.EvaluateAngles(c.in)
		if got != c.want {
			t.Errorf("%+v.evaluateAngles == %+v\n, want %+v\n", c.in, got, c.want)
		}
	}
}
