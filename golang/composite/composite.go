package composite

import "fmt"

type Node interface {
	Add(node Node)
	Ls(space int)
}

type File struct {
	name string
}

func (f *File) Add(node Node) {
	fmt.Println("不能添加子节点")
}

func (f *File) Ls(space int) {
	for i := 0; i < space; i++ {
		fmt.Print(" ")
	}
	fmt.Println(f.name)
}

type Folder struct {
	children []Node
	name     string
}

func (f *Folder) Add(node Node) {
	f.children = append(f.children, node)
}

func (f *Folder) Ls(space int) {
	for i := 0; i < space; i++ {
		fmt.Print(" ")
	}
	fmt.Println(f.name)
	space++
	for _, child := range f.children {
		child.Ls(space)
	}
}
