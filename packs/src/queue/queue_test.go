package Queue

import (
	"testing"
)

var size int
var [2]IQueue

func createLists(size int) {
  queuearray := &QueueArray{}
}

func deleteLists() {

}

func setupTest() func() {
	//before each test
	size = 10
	createLists(size)

	//after each test
	return func() {
		deleteLists()
	}
}

func TestAdd(t *testing.T) {
	defer setupTest()()
	for _, list := range lists {
		for i := 0; i < size; i++ {
			list.Add(i)
			if list.Size() != i+1 {
				t.Errorf("%T size is %d, but we expected it to be %d", list, list.Size(), i+1)
			}
		}
	}
}
