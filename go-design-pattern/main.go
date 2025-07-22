package main

import (
	"fmt"
	. "go-design-pattern/creational-patterns"
	s "go-design-pattern/structural-patterns"
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

	// 适配器模式
	// 创建一个 ThirdPartyLogger 实例
	thirdPartyLogger := &s.MyThirdPartyLogger{}

	// 使用适配器创建一个 Logger 实例
	logger = &s.ThirdPartyLoggerAdapter{Logger: thirdPartyLogger}

	// 使用 Logger 实例记录消息
	logger.Log("这是一条消息")

	// 桥接模式
	smsNotification := s.NewSMSNotification()     // 创建短信通知对象
	emailNotification := s.NewEmailNotification() // 创建邮件通知对象

	// 发送短信通知
	smsNotification.SendNotification("This is a SMS notification.")

	// 发送邮件通知
	emailNotification.SendNotification("This is an email notification.")

	// 组合模式
	// 创建文件系统
	root := &s.Directory{Name: "root"}
	documents := &s.Directory{Name: "documents"}
	pictures := &s.Directory{Name: "pictures"}
	music := &s.Directory{Name: "music"}
	file1 := &s.File{Name: "file1.txt"}
	file2 := &s.File{Name: "file2.txt"}
	file3 := &s.File{Name: "file3.txt"}

	root.AddChild(documents).
		AddChild(pictures).
		AddChild(music)
	documents.
		AddChild(file1).
		AddChild(file2)
	pictures.AddChild(file3)

	// 打印文件系统的层次结构
	root.List(0)

	// 装饰器模式
	toyota := s.Car{Brand: "Toyota", Price: 10000}
	decorator := s.ExtraPriceDecorator{ExtraPrice: 500}
	decoratedCar := decorator.DecoratePrice(toyota)
	fmt.Printf("%+v\n", decoratedCar)

	// 外观模式
	computer := s.NewComputerFacade()
	computer.Run()

	// 享元模式
	factory := &s.FlyweightFactory{Flyweights: make(map[string]s.Flyweight)}

	// 获取共享对象，如果工厂中不存在则创建
	flyweight1 := factory.GetFlyweight("flyweight1")
	flyweight2 := factory.GetFlyweight("flyweight2")
	flyweight3 := factory.GetFlyweight("flyweight3")
	flyweight4 := factory.GetFlyweight("flyweight4")
	flyweight5 := factory.GetFlyweight("flyweight5")

	// 调用共享对象的操作方法
	flyweight1.Operation("external state 1")
	flyweight2.Operation("external state 2")
	flyweight3.Operation("external state 3")
	flyweight4.Operation("external state 4")
	flyweight5.Operation("external state 5")

	// 代理模式
	proxy := &s.Proxy{}
	fmt.Println(proxy.Request())
}
