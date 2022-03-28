package _2_AbstractFactory_抽象工厂

import "testing"

func TestNewSimpleLunchFactory(t *testing.T) {
	factory := NewSimpleLunchFactory()
	food := factory.CreateFood()
	food.Cook()

	vegetable := factory.CreateVegetable()
	vegetable.Cook()
}
