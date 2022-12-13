package simple_factory

import "testing"

func TestType(t *testing.T) {
	api := NewAPI("Tank", 800)
	api.Show()
}

func TestType1(t *testing.T) {
	api := NewAPI("AirPlane", 800)
	api.Show()
}
