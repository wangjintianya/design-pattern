package state

import "fmt"

type State interface {
	SwitchOn(switcher *Switcher)
	SwitchOff(switcher *Switcher)
}

type Switcher struct {
	state State
}

func (s *Switcher) switchOn() {
	s.state.SwitchOn(s)
}

func (s *Switcher) switchOff() {
	s.state.SwitchOff(s)
}

type On struct {
}

func (o *On) SwitchOn(switcher *Switcher) {
	fmt.Println("WARN!!!通电状态无需再开")
}

func (o *On) SwitchOff(switcher *Switcher) {
	switcher.state = &Off{}
	fmt.Println("OK...灯灭")
}

type Off struct {
}

func (o *Off) SwitchOn(switcher *Switcher) {
	switcher.state = &On{}
	fmt.Println("OK...灯亮")
}

func (o *Off) SwitchOff(switcher *Switcher) {
	fmt.Println("WARN!!!断电状态无需再关")
}

func NewSwitcher() *Switcher {
	return &Switcher{
		state: &Off{},
	}
}
