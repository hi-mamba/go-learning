package main

// 这个提示先忽略,可以执行
import "golang.org/x/tour/pic"

/*


练习：slice

实现 `Pic`。它返回一个 slice 的长度 `dy`，和 slice 中每个元素的长度的 8 位无符号整数 `dx`。
当执行这个程序，它会将整数转换为灰度（好吧，蓝度）图片进行展示。

图片的实现已经完成。可能用到的函数包括 (x+y)/2 、 x*y 和 `x^y`（使用 math.Pow 计算最后的函数）。

（需要使用循环来分配 [][]uint8 中的每个 `[]uint8`。）

（使用 uint8(intValue) 在类型之间进行转换。）


*/


// TODO 目前这里下载包，由于被墙，暂时未解决
func Pic(dx, dy int) [][]uint8 {
	// 外层slice
	a := make([][]uint8, dy)
	for x := range a {
		// 里层slice
		b := make([]uint8, dx)
		for y := range b {
			// 给里层slice每个元素赋值
			b[y] = uint8(x*y - 1)
		}
		// 给外层slice每个元素赋值
		a[x] = b
	}
	return a
}
// https://blog.csdn.net/yjk13703623757/article/details/79619990
// https://blog.csdn.net/qq_27818541/article/details/54346106

func main() {
	// 这个提示先忽略,可以执行
	pic.Show(Pic)
}
