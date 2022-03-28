package _3_Facade_外观模式_门面模式

import "fmt"

// 类型
type CarModel struct{}

func NewCarModel() *CarModel {
	return &CarModel{}
}
func (c *CarModel) SetModel() {
	fmt.Println("CarModel - SetModel")
}

// 引擎
type CarEngine struct{}

func NewCarEngine() *CarEngine {
	return &CarEngine{}
}
func (c *CarEngine) SetEngine() {
	fmt.Println("CarEngine - SetEngine")
}

// 车身
type CarBody struct{}

func NewCarBody() *CarBody {
	return &CarBody{}
}
func (c *CarBody) SetBody() {
	fmt.Println("CarBody - SetBody")
}

// 门面模式-隐藏内部细节，提供一个门面访问
type CarFacade struct {
	model  CarModel
	engine CarEngine
	body   CarBody
}

func NewCarFacade() *CarFacade {
	return &CarFacade{
		model:  CarModel{},
		engine: CarEngine{},
		body:   CarBody{},
	}
}
func (c *CarFacade) CreateCompleteCar() {
	c.model.SetModel()
	c.engine.SetEngine()
	c.body.SetBody()
}
