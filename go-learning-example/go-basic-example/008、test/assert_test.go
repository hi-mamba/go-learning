package math

import (
	"math"
	"testing"
)

// 程序非常简洁，a 是 Abs 函数的输入参数，expect 是期望得到的执行结果，
// actual 是函数执行的实际结果，测试结果由 actual 和 expect 比较结果确定。
func TestAbs(t *testing.T) {
	var a, expect float64 = -10, 10

	actual := math.Abs(a)
	if actual != expect {
		t.Fatalf("a = %f, actual = %f, expected = %f", a, actual, expect)
	}
}