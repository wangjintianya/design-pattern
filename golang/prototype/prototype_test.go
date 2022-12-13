package prototype

import (
	"fmt"
	"testing"
)

var manager *Manager

func init() {
	manager = NewPrototypeManager()

	plane := EnemyPlane{
		X: 50,
	}

	manager.Set("p1", &plane)
}

func TestPrototype(t *testing.T) {
	plane1 := manager.Get("p1")
	plane2 := plane1.(*EnemyPlane)

	fmt.Println(plane2.X)
}
