package simple_factory

import (
	"fmt"
	"math/rand"
)

type Enemy struct {
	x, y int
}

type API interface {
	Show()
}

func NewAPI(enemyType string, screenWith int) API {
	x := rand.Intn(screenWith)
	switch enemyType {
	case "Tank":
		return &TankEnemy{
			&Enemy{
				x: x,
			},
		}
	case "AirPlane":
		return &AirPlaneEnemy{
			&Enemy{
				x: x,
			},
		}
	}
	return nil
}

type TankEnemy struct {
	*Enemy
}

func (t *TankEnemy) Show() {
	fmt.Printf("坦克出现的坐标：%d, %d \n", t.x, t.y)
	fmt.Println("坦克攻击玩家。。。")
}

type AirPlaneEnemy struct {
	*Enemy
}

func (a *AirPlaneEnemy) Show() {
	fmt.Printf("飞机出现的坐标：%d, %d \n", a.x, a.y)
	fmt.Println("飞机攻击玩家。。。")
}
