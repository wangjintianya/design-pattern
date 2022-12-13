package interpreter

import "testing"

func TestInterpreter(t *testing.T) {
	move := Move{500, 600}
	move.Interpret()
	sequence := Sequence{
		expressions: make([]Expression, 0),
	}
	sequence.expressions = append(sequence.expressions, NewLeftClick())
	sequence.expressions = append(sequence.expressions, &Delay{seconds: 1})
	repetition := Repetition{
		loopCount:  5,
		expression: &sequence,
	}
	repetition.Interpret()
	right := RightDown{}
	right.Interpret()
	delay := Delay{seconds: 7}
	delay.Interpret()
}
