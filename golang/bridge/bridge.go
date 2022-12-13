package bridge

import "fmt"

type Ruler interface {
	// Regularize 规则化的笔画走向
	Regularize()
}

type SquareRuler struct {
}

func (s *SquareRuler) Regularize() {
	fmt.Println("□")
}

type TriangleRuler struct {
}

func (s *TriangleRuler) Regularize() {
	fmt.Println("△")
}

type CircleRuler struct {
}

func (c *CircleRuler) Regularize() {
	fmt.Println("○")
}

type Drawable interface {
	Draw()
}

type BlackPen struct {
	ruler Ruler
}

func (b *BlackPen) Draw() {
	fmt.Print("黑")
	b.ruler.Regularize()
}

type RedPen struct {
	ruler Ruler
}

func (r *RedPen) Draw() {
	fmt.Print("红")
	r.ruler.Regularize()
}
