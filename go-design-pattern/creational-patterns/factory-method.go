package creationalpatterns

import "fmt"

// Logger 接口定义了一个日志记录器的方法
type Logger interface {
	Log(msg string)
}

// FileLogger 结构体实现了文件日志记录器
type FileLogger struct {
	filepath string
}

func (fl *FileLogger) Log(msg string) {
	// 实现文件日志记录的代码
	fmt.Println(msg)
}

// ConsoleLogger 结构体实现了控制台日志记录器
type ConsoleLogger struct{}

func (cl *ConsoleLogger) Log(msg string) {
	// 实现控制台日志输出的代码
	fmt.Println(msg)
}

// LoggerFactory 接口定义了一个工厂方法来创建日志记录器
type LoggerFactory interface {
	CreateLogger() Logger
}

// FileLoggerFactory 结构体实现了创建文件日志记录器的工厂方法
type FileLoggerFactory struct {
	Filepath string
}

func (flf *FileLoggerFactory) CreateLogger() Logger {
	return &FileLogger{filepath: flf.Filepath}
}

// ConsoleLoggerFactory 结构体实现了创建控制台日志记录器的工厂方法
type ConsoleLoggerFactory struct{}

func (clf *ConsoleLoggerFactory) CreateLogger() Logger {
	return &ConsoleLogger{}
}
