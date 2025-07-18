package creationalpatterns

import "sync"

// singleton 结构体定义了一个单例对象
type singleton struct {
	// ...
}

// instance 是一个包级别的变量，用于保存唯一的单例对象
var instance *singleton

// once 是一个用于确保初始化操作只执行一次的 sync.Once 对象
var once sync.Once

// GetInstance 方法返回唯一的单例对象
func GetInstance() *singleton {
	// 使用 sync.Once 确保初始化操作只执行一次
	once.Do(func() {
		instance = &singleton{
			// ...
		}
	})
	return instance
}
