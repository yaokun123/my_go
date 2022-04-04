package _2_decorator

import "fmt"

func ExampleDecorator() {
	var c Component = &ConcreteComponent{}
	c = WarpAddDecorator(c, 10)
	c = WarpMulDecorator(c, 8)
	res := c.Calc() // 8 * 10

	fmt.Printf("res %d\n", res)
	// Output:
	// res 80
}
