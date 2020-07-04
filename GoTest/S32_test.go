package GoTest

import "testing"

//testing.T 功能性测试
func TestGetArea(t *testing.T) {
	area := GetArea(40, 50)
	if area != 2000 {
		t.Error("测试失败")
	}
}
// go test -v

//testing.B  压力测试
func BenchmarkGetArea(t *testing.B) {

	for i := 0; i < t.N; i++ {
		GetArea(40, 50)
	}
}
//go test -bench=.