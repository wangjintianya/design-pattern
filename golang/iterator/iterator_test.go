package iterator

import (
	"fmt"
	"strconv"
	"testing"
)

func TestIterator(t *testing.T) {
	recorder := NewDrivingRecorder()
	for i := 0; i < 12; i++ {
		recorder.append("第" + strconv.Itoa(i) + "视频")
	}
	iterator := recorder.Iterator()
	for iterator.HasNext() {
		record := iterator.Next()
		if record == "第3视频" || record == "第6视频" {
			fmt.Println(record)
		}
	}
}
