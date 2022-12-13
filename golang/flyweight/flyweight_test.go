package flyweight

import "testing"

func TestFlyweight(t *testing.T) {
	factory := NewFactory()
	factory.GetDrawable("河流").Draw(10, 10)
	factory.GetDrawable("河流").Draw(10, 20)
	factory.GetDrawable("石路").Draw(10, 30)
	factory.GetDrawable("草坪").Draw(10, 40)
	factory.GetDrawable("草坪").Draw(10, 50)
	factory.GetDrawable("草坪").Draw(10, 60)
	factory.GetDrawable("草坪").Draw(10, 70)
	factory.GetDrawable("草坪").Draw(10, 80)
	factory.GetDrawable("房子").Draw(10, 80)
}
