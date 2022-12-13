package facade

import "testing"

func TestFacade(t *testing.T) {
	facade := NewFacade()
	facade.provideService()
}
