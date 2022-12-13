package facade

import "fmt"

// VegVendor 菜贩
type VegVendor struct {
}

func (v *VegVendor) Sell() {
	fmt.Println("卖菜")
}

// Chef 厨师
type Chef struct {
}

func (c *Chef) Cook() {
	fmt.Println("做菜")
}

// Waiter 服务员
type Waiter struct {
}

func (w *Waiter) Order() {
	fmt.Println("点菜")
}

func (w *Waiter) Serve() {
	fmt.Println("上菜")
}

// Cleaner 清理工
type Cleaner struct {
}

func (c *Cleaner) Clean() {
	fmt.Println("清理")
}

func (c *Cleaner) Wash() {
	fmt.Println("洗碗")
}

type Facade struct {
	VegVendor
	Chef
	Waiter
	Cleaner
}

func (f *Facade) provideService() {
	// 接待 点菜
	f.Order()
	// 厨师做菜
	f.Cook()
	// 上菜
	f.Serve()
	// 清理
	f.Clean()
	// 洗碗
	f.Wash()
}

func NewFacade() Facade {
	facade := Facade{}
	facade.Sell()
	return facade
}
