package sayHello

import (
	"fmt"
	"math"

	// 匿名导入：导入但不调用
	_ "math"
)

/*
*
首字母大写，可以被包外访问
首字母小写，外界无法访问
*/
func SayHello() {
	fmt.Println("Hello World!")
}
func calcFunc() {
	fmt.Println(math.MaxInt64)
}
