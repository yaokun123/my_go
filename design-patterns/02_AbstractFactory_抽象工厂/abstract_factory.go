package _2_AbstractFactory_抽象工厂

import "fmt"

type Lunch interface {
	Cook() // 做饭
}

type Rise struct{}

func (r *Rise) Cook() {
	fmt.Println("it is rise") // 做米饭
}

type Tomato struct{}

func (t *Tomato) Cook() {
	fmt.Println("ti is tomato") // 做西红柿
}

type LunchFactory interface { //抽象工厂，方便产品一族的扩展
	CreateFood() Lunch
	CreateVegetable() Lunch
}

type SimpleLunchFactory struct{} // 抽象工厂的一种实现（产品一族）
func (s *SimpleLunchFactory) CreateFood() Lunch {
	return &Rise{}
}
func (s *SimpleLunchFactory) CreateVegetable() Lunch {
	return &Tomato{}
}

func NewSimpleLunchFactory() LunchFactory {
	return &SimpleLunchFactory{}
}
