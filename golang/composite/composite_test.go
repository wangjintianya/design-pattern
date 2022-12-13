package composite

import "testing"

func TestComposite(t *testing.T) {
	driveD := Folder{
		name: "D盘",
	}
	doc := Folder{
		name: "文档",
	}
	doc.Add(&File{
		name: "简历.doc",
	})
	doc.Add(&File{
		name: "项目介绍.ppt",
	})
	driveD.Add(&doc)
	music := Folder{
		name: "音乐",
	}
	jay := Folder{
		name: "周杰伦",
	}
	jay.Add(&File{name: "双节棍.mp3"})
	jay.Add(&File{name: "告白气球.mp3"})
	jay.Add(&File{name: "听妈妈的话.mp3"})
	music.Add(&jay)
	jack := Folder{
		name: "张学友",
	}
	jack.Add(&File{name: "吻别.mp3"})
	jack.Add(&File{name: "一千个伤心的理由.mp3"})
	music.Add(&jack)
	driveD.Add(&music)

	driveD.Ls(0)
}
