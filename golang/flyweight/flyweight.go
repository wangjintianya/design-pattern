package flyweight

import "fmt"

type Drawable interface {
	Draw(x int, y int)
}

type Water struct {
	image string
}

func (w *Water) Draw(x int, y int) {
	fmt.Printf("在位置[%d, %d] 上绘制图片：[%s]\n", x, y, w.image)
}

func NewWater() *Water {
	image := "河流"
	fmt.Printf("从磁盘加载[%s]图片，耗时半秒。。。\n", image)
	return &Water{
		image: image,
	}
}

type Grass struct {
	image string
}

func (g *Grass) Draw(x int, y int) {
	fmt.Printf("在位置[%d, %d] 上绘制图片：[%s]\n", x, y, g.image)
}

func NewGrass() *Grass {
	image := "草坪"
	fmt.Printf("从磁盘加载[%s]图片，耗时半秒。。。\n", image)
	return &Grass{
		image: image,
	}
}

type Stone struct {
	image string
}

func (g *Stone) Draw(x int, y int) {
	fmt.Printf("在位置[%d, %d] 上绘制图片：[%s]\n", x, y, g.image)
}

func NewStone() *Stone {
	image := "石路"
	fmt.Printf("从磁盘加载[%s]图片，耗时半秒。。。\n", image)
	return &Stone{
		image: image,
	}
}

type House struct {
	image string
}

func (g *House) Draw(x int, y int) {
	fmt.Println("切换到顶层图片")
	fmt.Printf("在位置[%d, %d] 上绘制图片：[%s]\n", x, y, g.image)
}

func NewHouse() *House {
	image := "房子"
	fmt.Printf("从磁盘加载[%s]图片，耗时一秒。。。\n", image)
	return &House{
		image: image,
	}
}

type Factory struct {
	images map[string]Drawable
}

func (f *Factory) GetDrawable(image string) Drawable {
	_, ok := f.images[image]
	if !ok {
		switch image {
		case "河流":
			f.images[image] = NewWater()
		case "草坪":
			f.images[image] = NewGrass()
		case "石路":
			f.images[image] = NewStone()
		case "房子":
			f.images[image] = NewHouse()
		}
	}
	return f.images[image]
}

func NewFactory() *Factory {
	return &Factory{
		images: make(map[string]Drawable),
	}
}
