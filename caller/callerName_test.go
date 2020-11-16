package caller

import (
	"testing"
)

func TestController(t *testing.T) {
	Controller()
	//log.Println("function name : ", funcName)
}

func BenchmarkController(b *testing.B) {
	for n:=0; n<b.N;n++{
		Controller()
	}
}
