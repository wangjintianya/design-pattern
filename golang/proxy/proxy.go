package proxy

import (
	"fmt"
	"strings"
)

type Internet interface {
	Access(string)
}

type Modem struct {
}

func (m *Modem) Access(url string) {
	fmt.Printf("正在访问:%s\n", url)
}

type RouterProxy struct {
	modem     Modem
	blackList []string
}

func (r *RouterProxy) Access(url string) {
	for _, black := range r.blackList {
		if strings.Contains(url, black) {
			fmt.Printf("禁止访问:%s\n", url)
			return
		}
	}
	r.modem.Access(url)
}

func NewRouterProxy() *RouterProxy {
	return &RouterProxy{
		modem:     Modem{},
		blackList: []string{"电影", "游戏", "小说", "音乐"},
	}
}
