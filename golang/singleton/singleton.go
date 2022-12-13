package singleton

import "sync"

var (
	god  *God
	once sync.Once
)

type God struct {
}

func GetInstance() *God {
	once.Do(func() {
		god = &God{}
	})
	return god
}
