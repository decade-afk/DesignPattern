package creationalpatterns

// AbstractFactory 接口定义了一组工厂方法，用于创建一组相关的产品对象
type AbstractFactory interface {
	CreateProductA() ProductA
	CreateProductB() ProductB
}

// ProductA 接口定义了产品A的通用方法
type ProductA interface {
	Name() string
}

// ProductB 接口定义了产品B的通用方法
type ProductB interface {
	Name() string
}

// Factory1 结构体用于创建一组产品A和产品B的具体实现1
type Factory1 struct{}

func (f *Factory1) CreateProductA() ProductA {
	return &ProductA1{}
}

func (f *Factory1) CreateProductB() ProductB {
	return &ProductB1{}
}

// Factory2 结构体用于创建一组产品A和产品B的具体实现2
type Factory2 struct{}

func (f *Factory2) CreateProductA() ProductA {
	return &ProductA2{}
}

func (f *Factory2) CreateProductB() ProductB {
	return &ProductB2{}
}

// ProductA1 结构体是产品A的具体实现1
type ProductA1 struct{}

func (p *ProductA1) Name() string {
	return "Product A1"
}

// ProductA2 结构体是产品A的具体实现2
type ProductA2 struct{}

func (p *ProductA2) Name() string {
	return "Product A2"
}

// ProductB1 结构体是产品B的具体实现1
type ProductB1 struct{}

func (p *ProductB1) Name() string {
	return "Product B1"
}

// ProductB2 结构体是产品B的具体实现2
type ProductB2 struct{}

func (p *ProductB2) Name() string {
	return "Product B2"
}
