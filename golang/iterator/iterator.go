package iterator

import "fmt"

type DrivingRecorder struct {
	index   int
	records [10]string
}

func (d *DrivingRecorder) append(record string) {
	if d.index == 9 {
		d.index = 0
	} else {
		d.index++
	}
	d.records[d.index] = record
}

func (d *DrivingRecorder) display() {
	for i := 0; i < 10; i++ {
		fmt.Println(d.records[i])
	}
}

func (d *DrivingRecorder) displayInOrder() {
	for i, loopCount := d.index, 0; loopCount < 10; loopCount++ {
		fmt.Println(d.records[i])
		if i == 0 {
			i = 9
		} else {
			i--
		}
	}
}

func NewDrivingRecorder() *DrivingRecorder {
	return &DrivingRecorder{
		index: -1,
	}
}

type Iterator interface {
	HasNext() bool
	Next() string
}

type itr struct {
	cursor, loopCount int
	drivingRecorder   *DrivingRecorder
}

func (itr *itr) HasNext() bool {
	return itr.loopCount < 10
}

func (itr *itr) Next() string {
	i := itr.cursor
	if itr.cursor == 0 {
		itr.cursor = 9
	} else {
		itr.cursor--
	}
	itr.loopCount++
	return itr.drivingRecorder.records[i]
}

func (d *DrivingRecorder) Iterator() Iterator {
	return &itr{
		cursor:          d.index,
		drivingRecorder: d,
	}
}
