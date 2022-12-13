package adapter

import "testing"

func TestAdapter(t *testing.T) {
	adapter := Adapter{
		dualPinDevice: &TV{},
	}
	adapter.Electrify(1, 2, 3)
}
