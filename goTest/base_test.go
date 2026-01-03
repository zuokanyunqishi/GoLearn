package goTest

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAbc(t *testing.T) {

	Convey("测试", t, func() {
		So("a", ShouldEqual, "b")

	},
	)
}

func BenchmarkAbc(b *testing.B) {

	var c int
	for i := 0; i < b.N; i++ {
		c++
	}

}

func BenchmarkLogic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Logic to benchmark
	}
}
