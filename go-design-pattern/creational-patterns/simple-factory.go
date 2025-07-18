package creationalpatterns

// Shape 接口
type Shape interface {
	Draw() string
}

// Circle 结构体
type Circle struct{}

// Draw 方法实现
func (c *Circle) Draw() string {
	return "Circle"
}

// Rectangle 结构体
type Rectangle struct{}

// Draw 方法实现
func (r *Rectangle) Draw() string {
	return "Rectangle"
}

// 工厂函数
func NewShape(shapeType string) Shape {
	switch shapeType {
	case "circle":
		return new(Circle)
	case "rectangle":
		return new(Rectangle)
	default:
		return nil
	}
}
