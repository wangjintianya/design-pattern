package visitor

import (
	"fmt"
	"time"
)

type Acceptable interface {
	Accept(visitor Visitor)
}

type Visitor interface {
	visit(acceptable Acceptable)
}

type DiscountVisitor struct {
	billDate string
}

func GetDaysBetween2Date(format, date1Str, date2Str string) (int, error) {
	// 将字符串转化为Time格式
	date1, err := time.ParseInLocation(format, date1Str, time.Local)
	if err != nil {
		return 0, err
	}
	// 将字符串转化为Time格式
	date2, err := time.ParseInLocation(format, date2Str, time.Local)
	if err != nil {
		return 0, err
	}
	//计算相差天数
	return int(date1.Sub(date2).Hours() / 24), nil
}

func (d *DiscountVisitor) visit(acceptable Acceptable) {
	switch a := acceptable.(type) {
	case *Candy:
		fmt.Printf("=====糖果【%s】打折后价格=====", a.name)
		days, _ := GetDaysBetween2Date("2006-01-02", d.billDate, a.produceDate)
		rate := 0.0
		if days > 180 {
			fmt.Println("超过半年过期糖果，请勿食用！")
		} else {
			rate = 0.9
		}
		fmt.Printf("￥ %f\n", rate*a.price)
	case *Fruit:
		fmt.Printf("=====水果【%s】打折后价格=====", a.name)
		days, _ := GetDaysBetween2Date("2006-01-02", d.billDate, a.produceDate)
		rate := 0.0
		if days > 7 {
			fmt.Println("￥0.00元（超过一周过期水果，请勿食用！）")
		} else if days > 3 {
			rate = 0.5
		} else {
			rate = 0.9
		}
		fmt.Printf("￥ %f\n", rate*a.price*a.weight)
	case *Wine:
		fmt.Printf("=====酒品【%s】无折扣价格=====", a.name)
		fmt.Printf("￥ %f\n", a.price)
	}
}

type Product struct {
	name        string
	price       float64
	produceDate string
}

type Candy struct {
	Product
}

func (c *Candy) Accept(visitor Visitor) {
	visitor.visit(c)
}

func NewCandy(name string, price float64, produceDate string) *Candy {
	return &Candy{
		Product{
			name:        name,
			price:       price,
			produceDate: produceDate,
		},
	}
}

type Wine struct {
	Product
}

func NewWine(name string, price float64, produceDate string) *Wine {
	return &Wine{
		Product{
			name:        name,
			price:       price,
			produceDate: produceDate,
		},
	}
}

func (w *Wine) Accept(visitor Visitor) {
	visitor.visit(w)
}

type Fruit struct {
	Product
	weight float64
}

func NewFruit(name string, price float64, produceDate string, weight float64) *Fruit {
	return &Fruit{
		Product{
			name:        name,
			price:       price,
			produceDate: produceDate,
		},
		weight,
	}
}

func (f *Fruit) Accept(visitor Visitor) {
	visitor.visit(f)
}
