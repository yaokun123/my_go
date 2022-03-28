package _1_Factory_工厂模式

import (
	"fmt"
	"testing"
)

func TestNewRestaurant(t *testing.T) {
	// 简单工厂
	fmt.Println("===============简单工厂的测试===============")
	NewRestaurant("d").GetFood()
	NewRestaurant("q").GetFood()

	// 工厂
	fmt.Println("===============工厂模式的测试===============")
	NewDonglaishun().GetFood()
	NewQingfeng().GetFood()
}
