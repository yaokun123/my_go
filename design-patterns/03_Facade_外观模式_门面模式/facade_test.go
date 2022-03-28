package _3_Facade_外观模式_门面模式

import "testing"

func TestNewCarFacade(t *testing.T) {
	//NewCarFacade().CreateCompleteCar()
	car := CarFacade{}
	car.CreateCompleteCar()
}
