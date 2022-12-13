package adapter

import "fmt"

type TriplePin interface { // 三线插口接口
	Electrify(l int, n int, e int) // l 火线 ， n 零线， e 地线
}

type DualPin interface { // 两线插口接口
	Electrify(l int, n int) // l 火线 ， n 零线
}

type TV struct {
}

func (t *TV) Electrify(l int, n int) {
	fmt.Printf("火线通电：%d\n", l)
	fmt.Printf("零线通电：%d\n", n)
}

type Adapter struct {
	dualPinDevice DualPin
}

func (a *Adapter) Electrify(l int, n int, e int) {
	a.dualPinDevice.Electrify(l, n)
}
