package main

import (
	"fmt"
	. "go-design-pattern/creational-patterns"
)

func main() {
	// 简单工厂模式
	circle := NewShape("circle")
	rectangle := NewShape("rectangle")
	fmt.Println(circle.Draw())    // 输出: Circle
	fmt.Println(rectangle.Draw()) // 输出: Rectangle

	// 工厂方法模式
	var loggerFactory LoggerFactory
	loggerFactory = &FileLoggerFactory{Filepath: "/var/log/myapp.log"}
	logger := loggerFactory.CreateLogger()
	logger.Log("An error occurred")

	loggerFactory = &ConsoleLoggerFactory{}
	logger = loggerFactory.CreateLogger()
	logger.Log("An error occurred")

	// 抽象工厂模式
	// 创建一个工厂1
	factory1 := &Factory1{}

	// 创建产品A1和产品B1
	productA1 := factory1.CreateProductA()
	productB1 := factory1.CreateProductB()

	fmt.Println(productA1.Name()) // Output: Product A1
	fmt.Println(productB1.Name()) // Output: Product B1

	// 创建一个工厂2
	factory2 := &Factory2{}

	// 创建产品A2和产品B2
	productA2 := factory2.CreateProductA()
	productB2 := factory2.CreateProductB()

	fmt.Println(productA2.Name()) // Output: Product A2
	fmt.Println(productB2.Name()) // Output: Product B2

	// 单例模式

	// 建造者模式
	// 创建一个 PersonBuilder 实例
	builder := &PersonBuilder{}

	// 设置人的信息
	person := builder.SetName("Tom").
		SetAge(18).
		SetGender("Male").
		SetAddress("123 Main St").
		Build()

	// 调用人的说话和睡觉方法
	person.Speak("Hello, World!")
	person.Sleep()

	// 原型模式
	// 创建原型对象
	prototype := &ConcretePrototype{Name: "prototype"}

	// 使用原型对象创建新对象
	newObject := prototype.Clone()

	// 打印新对象名称
	fmt.Println(newObject.(*ConcretePrototype).Name)

}
