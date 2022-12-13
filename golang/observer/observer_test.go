package observer

import "testing"

func TestObserver(t *testing.T) {
	shop := Shop{}
	phoneFan := NewPhoneFans("果粉唐僧", &shop)
	chopper := NewHandChopper("剁手族八戒", &shop)
	shop.register(chopper)
	shop.register(phoneFan)

	shop.setProduct("猪肉炖粉条")
	shop.setProduct("水果手机【爱疯叉】")
}
