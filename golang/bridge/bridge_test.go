package bridge

import "testing"

func TestBridge(t *testing.T) {
	// 黑色的所有图形
	pen := BlackPen{
		&CircleRuler{},
	}
	pen.Draw()
	pen = BlackPen{
		&TriangleRuler{},
	}
	pen.Draw()
	pen = BlackPen{
		&SquareRuler{},
	}
	pen.Draw()
	// 红色的所有图形
	pen1 := RedPen{
		&CircleRuler{},
	}
	pen1.Draw()
	pen1 = RedPen{
		&TriangleRuler{},
	}
	pen1.Draw()
	pen1 = RedPen{
		&SquareRuler{},
	}
	pen1.Draw()
}
