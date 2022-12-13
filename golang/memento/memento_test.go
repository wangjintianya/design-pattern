package memento

import "testing"

func TestMemento(t *testing.T) {
	doc := Doc{title: "《AI的觉醒》"}
	editor := NewEditor(doc)
	editor.append("第一章 混沌初开")
	editor.append("\n  正文2000字……")
	editor.append("\n第二章 荒漠之花\n  正文3000字……")

	editor.delete()
	editor.undo()
}
