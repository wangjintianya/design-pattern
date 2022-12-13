package visitor

import "testing"

func TestVisitor(t *testing.T) {
	var products []Acceptable
	products = append(products, NewCandy("小黑兔糖果", 2.0, "2022-01-01"))
	products = append(products, NewWine("猫泰白酒", 300, "2019-03-14"))
	products = append(products, NewFruit("草莓", 28, "2022-11-02", 3.0))

	visitor := DiscountVisitor{
		billDate: "2022-11-08",
	}

	for _, product := range products {
		product.Accept(&visitor)
	}
}
