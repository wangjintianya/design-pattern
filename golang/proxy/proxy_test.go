package proxy

import "testing"

func TestProxy(t *testing.T) {
	proxy := NewRouterProxy()
	proxy.Access("http://www.电影.com")
	proxy.Access("http://www.游戏.com")
	proxy.Access("http://www.读书.com")
}
