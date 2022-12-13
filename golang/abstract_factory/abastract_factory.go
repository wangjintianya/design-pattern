package abstract_factory

import "fmt"

type Unit interface {
	Show()
	Attack()
}

type Marine struct {
	attack  int // 攻击力
	defense int // 防御力
	health  int // 血量
	x, y    int // 坐标
}

func (m *Marine) Show() {
	fmt.Printf("士兵出现在坐标：[%d, %d]\n", m.x, m.y)
}

func (m *Marine) Attack() {
	fmt.Printf("士兵用机关枪射击，攻击力：%d\n", m.attack)
}

type Tank struct {
	attack  int // 攻击力
	defense int // 防御力
	health  int // 血量
	x, y    int // 坐标
}

func (t *Tank) Show() {
	fmt.Printf("坦克出现在坐标：[%d, %d]\n", t.x, t.y)
}

func (t *Tank) Attack() {
	fmt.Printf("坦克用炮轰击，攻击力：%d\n", t.attack)
}

type BattleShip struct {
	attack  int // 攻击力
	defense int // 防御力
	health  int // 血量
	x, y    int // 坐标
}

func (b *BattleShip) Show() {
	fmt.Printf("战舰出现在坐标：[%d, %d]\n", b.x, b.y)
}

func (b *BattleShip) Attack() {
	fmt.Printf("战舰用激光炮打击，攻击力：%d\n", b.attack)
}

type Roach struct {
	attack  int // 攻击力
	defense int // 防御力
	health  int // 血量
	x, y    int // 坐标
}

func (r *Roach) Show() {
	fmt.Printf("蟑螂兵出现在坐标：[%d, %d]\n", r.x, r.y)
}

func (r *Roach) Attack() {
	fmt.Printf("蟑螂兵用爪子挠，攻击力：%d\n", r.attack)
}

type Spitter struct {
	attack  int // 攻击力
	defense int // 防御力
	health  int // 血量
	x, y    int // 坐标
}

func (s *Spitter) Show() {
	fmt.Printf("口水出现在坐标：[%d, %d]\n", s.x, s.y)
}

func (s *Spitter) Attack() {
	fmt.Printf("口水兵用毒液喷射，攻击力：%d\n", s.attack)
}

type Mammoth struct {
	attack  int // 攻击力
	defense int // 防御力
	health  int // 血量
	x, y    int // 坐标
}

func (m *Mammoth) Show() {
	fmt.Printf("猛犸巨兽出现在坐标：[%d, %d]\n", m.x, m.y)
}

func (m *Mammoth) Attack() {
	fmt.Printf("猛犸巨兽用獠牙顶，攻击力：%d\n", m.attack)
}

type AbstractFactory interface {
	createLowClass() Unit
	createMidClass() Unit
	createHighClass() Unit
}

type HumanFactory struct {
	x, y int
}

func (h *HumanFactory) createLowClass() Unit {
	unit := Marine{
		attack:  6,
		defense: 5,
		health:  40,
		x:       h.x,
		y:       h.y,
	}
	fmt.Println("制作海军陆战队员成功。")
	return &unit
}

func (h *HumanFactory) createMidClass() Unit {
	unit := Tank{
		attack:  25,
		defense: 100,
		health:  150,
		x:       h.x,
		y:       h.y,
	}
	fmt.Println("制作变形坦克成功。")
	return &unit
}

func (h *HumanFactory) createHighClass() Unit {
	unit := BattleShip{
		attack:  35,
		defense: 200,
		health:  500,
		x:       h.x,
		y:       h.y,
	}
	fmt.Println("制作巨型战舰成功。")
	return &unit
}

type AlienFactory struct {
	x, y int
}

func (a *AlienFactory) createLowClass() Unit {
	unit := Roach{
		attack:  5,
		defense: 2,
		health:  35,
		x:       a.x,
		y:       a.y,
	}
	fmt.Println("制作蟑螂兵成功。")
	return &unit
}

func (a *AlienFactory) createMidClass() Unit {
	unit := Spitter{
		attack:  10,
		defense: 8,
		health:  80,
		x:       a.x,
		y:       a.y,
	}
	fmt.Println("制作毒液兵成功。")
	return &unit
}

func (a *AlienFactory) createHighClass() Unit {
	unit := Mammoth{
		attack:  10,
		defense: 8,
		health:  80,
		x:       a.x,
		y:       a.y,
	}
	fmt.Println("制作猛犸巨兽成功。")
	return &unit
}
