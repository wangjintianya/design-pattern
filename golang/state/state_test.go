package state

import "testing"

func TestState(t *testing.T) {
	switcher := NewSwitcher()
	switcher.switchOn()
	switcher.switchOn()
	switcher.switchOff()
	switcher.switchOff()
}
