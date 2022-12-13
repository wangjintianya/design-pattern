package decorator

import "testing"

func TestDecorator(t *testing.T) {
	girl := Girl{}
	makeup := FoundationMakeup{
		shower: &girl,
	}
	lipstick := Lipstick{
		shower: &makeup,
	}
	lipstick.Show()
}
