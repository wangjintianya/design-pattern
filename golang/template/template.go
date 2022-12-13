package template

import "fmt"

type Standard interface {
	Kickoff()
}

type template struct {
	im implement
}

type implement interface {
	analyse()
	design()
	develop()
	test() bool
	release()
}

func (t *template) Kickoff() {
	t.im.analyse()
	t.im.design()
	t.im.develop()
	for !t.im.test() {
		t.im.develop()
	}
	t.im.release()
}

func NewTemplate(im implement) *template {
	return &template{
		im: im,
	}
}

type Developer struct {
	*template
}

func (d *Developer) analyse() {
	fmt.Println("developer analyse...")
}
func (d *Developer) design() {
	fmt.Println("developer design...")
}
func (d *Developer) develop() {
	fmt.Println("developer develop...")
}
func (d *Developer) test() bool {
	fmt.Println("developer test...")
	return true
}
func (d *Developer) release() {
	fmt.Println("developer release...")
}

func NewDeveloper() *Developer {
	developer := &Developer{}
	newTemplate := NewTemplate(developer)
	return &Developer{
		template: newTemplate,
	}
}
