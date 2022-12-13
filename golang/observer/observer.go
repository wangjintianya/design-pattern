package observer

import (
	"fmt"
	"strings"
)

type Shop struct {
	product string
	buyers  []Observed
}

func (s *Shop) register(buyer Observed) {
	s.buyers = append(s.buyers, buyer)
}

func (s *Shop) setProduct(product string) {
	s.product = product
	s.notifyBuyers()
}

func (s *Shop) notifyBuyers() {
	for _, buyer := range s.buyers {
		buyer.inform()
	}
}

type Observed interface {
	inform()
}

type Buyer struct {
	name string
	shop *Shop
}

type HandChopper struct {
	Buyer
}

func (h *HandChopper) inform() {
	fmt.Println(h.name)
	fmt.Printf("购买%s\n", h.shop.product)
}

func NewHandChopper(name string, shop *Shop) *HandChopper {
	return &HandChopper{
		Buyer{
			name: name,
			shop: shop,
		},
	}
}

type PhoneFans struct {
	Buyer
}

func (p *PhoneFans) inform() {
	product := p.shop.product
	if strings.Contains(product, "水果手机") {
		fmt.Println(p.name)
		fmt.Printf("购买%s\n", product)
	}
}

func NewPhoneFans(name string, shop *Shop) *PhoneFans {
	return &PhoneFans{
		Buyer{
			name: name,
			shop: shop,
		},
	}
}
