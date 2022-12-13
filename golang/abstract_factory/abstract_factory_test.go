package abstract_factory

import (
	"fmt"
	"testing"
)

func TestGame(t *testing.T) {
	fmt.Println("游戏开始。。。")
	fmt.Println("双方挖矿攒钱。。。")
	//第一位玩家选择了地球人族
	fmt.Println("工人建造人族工厂。。。")
	factory := HumanFactory{
		x: 10,
		y: 10,
	}
	marine := factory.createLowClass()
	marine.Show()
	tank := factory.createMidClass()
	tank.Show()
	battleShip := factory.createHighClass()
	battleShip.Show()

	//另一位玩家选择了外星族
	fmt.Println("工蜂建造外星虫族工厂。。。")
	factoryNew := AlienFactory{
		x: 10,
		y: 10,
	}
	roach := factoryNew.createLowClass()
	roach.Show()
	spitter := factoryNew.createMidClass()
	spitter.Show()
	mammoth := factoryNew.createHighClass()
	mammoth.Show()

	fmt.Println("两族开始大混战。。。")
	marine.Attack()
	roach.Attack()
	spitter.Attack()
	tank.Attack()
	mammoth.Attack()
	battleShip.Attack()

}
