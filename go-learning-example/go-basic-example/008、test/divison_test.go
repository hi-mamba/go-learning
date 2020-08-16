package math

import (
	"errors"
	"testing"
)

func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}

	return a / b, nil
}

// 定义了三个变量，分别是 a、b、expect，对应被除数、除数和期望结果。
//用例通过对比 Division 的实际结果 actual 与期望结果 expect 确认测试是否成功。
//还有就是，Division 返回的 error 也要检查，因为这里期待的正常运行结果，只要有错即可认定测试失败。

func TestDivision(t *testing.T) {
	var a, b, expect float64 = 10, 5, 2

	actual, err := Division(a, b)
	if err != nil {
		t.Errorf("a = %f, b = %f, expect = %f, err %v", a, b, expect, err)
		return
	}

	if actual != expect {
		t.Errorf("a = %f, b = %f, expect = %f, actual = %f", a, b, expect, actual)
	}
}

//同样是首先定义了三个变量，a、b 和 expectErrString，a、b 含义与之前相同，
//expectErrString 为预期提示的错误信息。除数 b 设置为 0 ，
//主要是为了测试 Division 函数是否能按预期返回错误，所以我们并不关心计算结果。
//测试成功与否，通过比较实际的返回 error 与 expectErrString 确定。

func TestDivisionZero(t *testing.T) {
	var a, b float64 = 10, 0
	var expectedErrString = "division by zero"

	_, err := Division(a, b)
	if err.Error() != expectedErrString {
		t.Errorf("a = %f, b = %f, err %v, expect err %s", a, b, err, expectedErrString)
		return
	}
}
