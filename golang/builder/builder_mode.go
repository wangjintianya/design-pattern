package builder

import "fmt"

type Builder interface {
	BuildBasement()
	BuildWall()
	BuildRoof()
	GetBuilding() Building
}

type HouseBuilder struct {
	house Building
}

func (h *HouseBuilder) BuildBasement() {
	fmt.Println("挖地基，部署管道、线缆，水泥加固，搭建围墙、花园。")
	h.house.setBasement("╬╬╬╬╬╬╬╬\n")
}

func (h *HouseBuilder) BuildWall() {
	fmt.Println("搭建木质框架，石膏板封墙并粉饰内外墙。")
	h.house.setWall("｜田｜田 田|\n")
}
func (h *HouseBuilder) BuildRoof() {
	fmt.Println("建造木质屋顶、阁楼，安装烟囱，做好防水。")
	h.house.setRoof("╱◥███◣\n")
}
func (h *HouseBuilder) GetBuilding() Building {
	return h.house
}

type ApartmentBuilder struct {
	apartment Building
}

func (a *ApartmentBuilder) BuildBasement() {
	fmt.Println("深挖地基，修建地下车库，部署管道、线缆、风道。")
	a.apartment.setBasement("╚═════════╝\n")
}

func (a *ApartmentBuilder) BuildWall() {
	fmt.Println("搭建多层建筑框架，建造电梯井，钢筋混凝土浇灌。")
	for i := 0; i < 8; i++ {
		a.apartment.setWall("║ □ □ □ □ ║\n")
	}
}
func (a *ApartmentBuilder) BuildRoof() {
	fmt.Println("封顶，部署通风井，做防水层，保温层。")
	a.apartment.setRoof("╔═════════╗\n")
}
func (a *ApartmentBuilder) GetBuilding() Building {
	return a.apartment
}
