package decorator

import "fmt"

type Showable interface {
	Show()
}

type Girl struct {
}

func (g *Girl) Show() {
	fmt.Print("女孩的素颜")
}

type FoundationMakeup struct {
	shower Showable
}

func (f *FoundationMakeup) Show() {
	fmt.Print("打粉底(")
	f.shower.Show()
	fmt.Print(")")
}

type Lipstick struct {
	shower Showable
}

func (l *Lipstick) Show() {
	fmt.Print("涂口红(")
	l.shower.Show()
	fmt.Print(")")
}
