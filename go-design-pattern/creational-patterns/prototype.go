package creationalpatterns

// 定义接口
type Prototype interface {
	Clone() Prototype
}

// 定义结构体
type ConcretePrototype struct {
	Name string
}

// 实现 Clone 方法
func (p *ConcretePrototype) Clone() Prototype {
	newPrototype := *p
	return &newPrototype
}
