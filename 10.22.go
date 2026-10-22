#声明一个名为 i 的整数变量，其值为 20。将 i 赋值给一个名为 f 的浮点变量。打印出 i 和 f 的值。
package main

import "fmt"

func main() {
	var i int = 20
	var f float64
	f = float64(i)
	fmt.Printf("i = %d,f = %f\n",i,f)
}
#声明一个名为 value 的常量，它可以被赋值给整数和浮点数变量。将其赋值给 i 和 f，打印。

package main

import "fmt"

func main(){
	const value = 100
	var i int
	var f float64
	i = value
	f = value
	fmt.Printf("i = %d,f = %f\n",i,f)
}
#编写一个包含 3 个变量的程序，其中一个名为 b（byte 类型），一个名为 smallI（int32 类型），一个名为 bigI（uint64 类型）。给每个变量分配其类型的最大法定值；然后，给每个变量加 1。打印出它们的值。

package main

import(
	"fmt"
	"math"
)

func main(){
	var b byte = math.MaxUint8
	var smallI = math.MaxInt32
	var bigI - math.MaxUint64

	b += 1
	smallI += 1
	bigI += 1

	fmt.Printf("b = %d\n",b)
	fmt.Printf("smallI = %d\n",smallI)
	fmt.Printf("bigI = %d\n",bigI)
}

