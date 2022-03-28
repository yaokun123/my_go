package main

import "fmt"

//1、类型断言
/**
类型断言就是将接口类型的值(x)，转换成类型(T)。格式为：
x.(T)
v := x.(T)
v,ok := x.(T)

类型断言的必要条件就是x是接口类型，非接口类型的x不能做类型断言
var i int = 10
v := i.(int)	//错误 i不是接口类型

T可以是非接口类型，如果想断言合法，则T应该实现x的接口
T也可以是接口，则x的动态类型也应该实现接口T
var x interface{} = 7	//x的动态类型为int，值为7
i := x.(int)			//i的类型为int，值为7

type I interface{m()}
var y I
s := y.(string)			//非法：string没有实现接口I(missing method m)
r := y.(io.Reader)		//y如果实现了接口io.Reader和I的情况下，r的类型则为io.Reader

类型断言如果非法，运行时就会出现错误，为了避免这种错误，可以使用一下语法：
v,ok := x.(T)
ok代表类型断言是否合法，如果非法，ok则为false，这样就不会出现panic了
*/

//2、类型切换
/**
类型切换用来比较类型而不是对值进行比较
type switch它用于检测的是值x的类型T是否匹配某个类型.
格式如下，类似类型断言，但是括号内的不是某个具体的类型，而是单词type:
switch x.(type){

}

type switch语句中可以有一个简写的变量声明，这种情况下，等价于这个变量声明在每个case clause 隐式代码块的开始位置。
如果case clause只列出了一个类型，则变量的类型就是这个类型，否则就是原始值的类型
假设下面的例子中的x的类型为x interface{}:

switch i := x.(type){
	case nil:
		fmt.Println("x is nil")		//i的类型是x的类型(interface{})
	case int:
		fmt.Println("x is int")
	case float64:
		fmt.Println("x is float64")
	case bool,string:
		fmt.Println("x is bool or string")
	default:
		fmt.Println("don't know the type")
}
*/

//3、打印类型
/**
fmt.Printf("%T",x)	//打印x的类型
*/

func main() {
	var x interface{}
	x = 12

	var y interface{}
	y = "test"
	fmt.Printf("x is %T\n", x)
	fmt.Printf("y is %T\n", y)

	x1 := x.(int)
	fmt.Printf("x1 is %T\n", x1)
	y1 := y.(string)
	fmt.Printf("y1 is %T\n", y1)

	//以下错误
	/*x2 := x.(string)
	fmt.Printf("x2 is %T\n",x2)
	y2 := y.(int)
	fmt.Printf("y2 is %T\n",y2)*/
}
