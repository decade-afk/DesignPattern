package structuralpatterns

// Subject 接口定义了需要被代理的对象的方法
type Subject interface {
	Request() string
}

// RealSubject 是需要被代理的对象
type RealSubject struct{}

func (r *RealSubject) Request() string {
	return "RealSubject: 处理请求"
}

// Proxy 是代理对象，它包含了一个指向 RealSubject 的引用
type Proxy struct {
	realSubject *RealSubject
}

func (p *Proxy) Request() string {
	// 在这里可以进行一些额外的操作，例如鉴权、日志等
	result := "Proxy: 转发请求到 RealSubject\n"
	if p.realSubject == nil {
		p.realSubject = &RealSubject{}
	}
	result += " " + p.realSubject.Request()
	return result
}
