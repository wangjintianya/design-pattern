package builder

import (
	"fmt"
	"testing"
)

func TestBuilderMode(t *testing.T) {
	//招工，建别墅。
	builder := HouseBuilder{}
	//交给工程总监
	director := Director{
		builder: &builder,
	}
	building := director.direct()
	fmt.Println(building.printInfo())
	//替换施工方，建公寓。
	director.builder = &ApartmentBuilder{}
	building = director.direct()
	fmt.Println(building.printInfo())
}
