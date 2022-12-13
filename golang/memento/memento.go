package memento

import "fmt"

type History struct {
	body string
}

type Doc struct {
	title string
	body  string
}

func (d *Doc) createHistory() *History {
	return &History{body: d.body}
}

func (d *Doc) restoreHistory(history *History) {
	d.body = history.body
}

type Editor struct {
	doc             Doc
	histories       []*History
	historyPosition int
}

func (e *Editor) show() {
	fmt.Println(e.doc.body)
	fmt.Println("文章结束>>>")
}

func (e *Editor) delete() {
	fmt.Println("<<<删除操作")
	e.doc.body = ""
	e.show()
}

func (e *Editor) save() {
	fmt.Println("<<<存盘操作")
}

func (e *Editor) append(txt string) {
	fmt.Println("<<<插入操作")
	e.backup()
	e.doc.body = e.doc.body + txt
	e.show()
}
func (e *Editor) backup() {
	e.histories = append(e.histories, e.doc.createHistory())
	e.historyPosition++
}

func (e *Editor) undo() {
	fmt.Println("<<<撤销操作")
	e.doc.restoreHistory(e.histories[e.historyPosition])
	e.historyPosition--
	e.show()
}

func NewEditor(doc Doc) *Editor {
	e := &Editor{
		doc:             doc,
		histories:       make([]*History, 0),
		historyPosition: -1,
	}
	fmt.Println("<<<打开文档" + doc.title)
	e.backup()
	e.show()

	return e
}
