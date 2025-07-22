package structuralpatterns

import "fmt"

// Flyweight 接口定义享元对象的方法
type Flyweight interface {
	Operation(extrinsicState string)
}

// ConcreteFlyweight 是具体的享元类，包含内部状态和外部状态
type ConcreteFlyweight struct {
	intrinsicState string // 内部状态
}

// Operation 实现 Flyweight 接口中定义的操作方法
func (f *ConcreteFlyweight) Operation(extrinsicState string) {
	fmt.Printf("ConcreteFlyweight: intrinsic state is %s, extrinsic state is %s\n", f.intrinsicState, extrinsicState)
}

// FlyweightFactory 是享元工厂类，负责创建和管理共享对象
type FlyweightFactory struct {
	Flyweights map[string]Flyweight
}

// GetFlyweight 从享元工厂中获取共享对象
func (factory *FlyweightFactory) GetFlyweight(key string) Flyweight {
	if flyweight, ok := factory.Flyweights[key]; ok {
		return flyweight
	}
	flyweight := &ConcreteFlyweight{intrinsicState: key}
	factory.Flyweights[key] = flyweight
	return flyweight
}
