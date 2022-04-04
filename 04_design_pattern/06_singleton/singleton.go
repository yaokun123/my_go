package _6_singleton

import "sync"

/**
使用懒汉式的单例模式，使用双重检查加锁保证线程安全

还可以使用饿汉式
*/

//Singleton 是单例模式类
type Singleton struct{}

var singleton *Singleton
var once sync.Once

//GetInstance 用于获取单例模式对象
func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})

	return singleton
}
