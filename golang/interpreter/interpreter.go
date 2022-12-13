package interpreter

import (
	"fmt"
	"time"
)

type Expression interface {
	Interpret()
}

type Move struct {
	x, y int
}

func (m *Move) Interpret() {
	fmt.Printf("移动鼠标:[%d, %d]\n", m.x, m.y)
}

type Delay struct {
	seconds int
}

func (d *Delay) Interpret() {
	fmt.Printf("系统延时%d秒\n", d.seconds)
	time.Sleep(time.Duration(d.seconds) * time.Second)
}

type LeftDown struct {
}

func (l *LeftDown) Interpret() {
	fmt.Println("按下鼠标:左键")
}

type RightDown struct {
}

func (r *RightDown) Interpret() {
	fmt.Println("按下鼠标:右键")
}

type LeftUp struct {
}

func (l *LeftUp) Interpret() {
	fmt.Println("松开鼠标:左键")
}

type LeftClick struct {
	leftDown LeftDown
	leftUp   LeftUp
}

func (l *LeftClick) Interpret() {
	l.leftDown.Interpret()
	l.leftUp.Interpret()
}

func NewLeftClick() *LeftClick {
	return &LeftClick{
		leftDown: LeftDown{},
		leftUp:   LeftUp{},
	}
}

type Repetition struct {
	loopCount  int
	expression Expression
}

func (r *Repetition) Interpret() {
	for r.loopCount > 0 {
		r.expression.Interpret()
		r.loopCount--
	}
}

type Sequence struct {
	expressions []Expression
}

func (s *Sequence) Interpret() {
	for _, expression := range s.expressions {
		expression.Interpret()
	}
}
