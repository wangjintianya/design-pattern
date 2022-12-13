package builder

import "fmt"

type Director struct {
	builder Builder
}

func (d *Director) direct() Building {
	fmt.Println("=====工程项目启动=====")
	// 第一步 打好地基
	d.builder.BuildBasement()
	// 第二步 建造框架 墙体
	d.builder.BuildWall()
	// 第三步 封顶
	d.builder.BuildRoof()
	fmt.Println("=====工程项目竣工=====")

	return d.builder.GetBuilding()
}
