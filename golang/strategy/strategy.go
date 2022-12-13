package strategy

type Strategy interface {
	Calculate(a int, b int) int
}

type Addition struct {
}

func (adt *Addition) Calculate(a int, b int) int {
	return a + b
}

type Subtraction struct {
}

func (s *Subtraction) Calculate(a int, b int) int {
	return a - b
}

type Calculator struct {
	strategy Strategy
}

func (c *Calculator) GetResult(a int, b int) int {
	return c.strategy.Calculate(a, b)
}
