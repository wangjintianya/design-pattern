package strategy

import (
	"fmt"
	"testing"
)

func TestStrategy(t *testing.T) {
	calculator := Calculator{}
	calculator.strategy = &Addition{}
	fmt.Println(calculator.GetResult(1, 2))
	calculator.strategy = &Subtraction{}
	fmt.Println(calculator.GetResult(1, 2))
}
