package structuralpatterns

import "fmt"

// 实现发送消息的接口
type Sender interface {
	Send(message string)
}

// 短信发送器
type SMS struct{}

func (s *SMS) Send(message string) {
	fmt.Println("Sending SMS:", message)
}

// 邮件发送器
type Email struct{}

func (e *Email) Send(message string) {
	fmt.Println("Sending Email:", message)
}

// 抽象通知类
type Notification struct {
	sender Sender // 保存发送器接口的实例
}

func (n *Notification) SendNotification(message string) {
	n.sender.Send(message)
}

// 短信通知类
type SMSNotification struct {
	Notification // 组合抽象通知类
}

func NewSMSNotification() *SMSNotification {
	return &SMSNotification{Notification{&SMS{}}}
}

// 邮件通知类
type EmailNotification struct {
	Notification // 组合抽象通知类
}

func NewEmailNotification() *EmailNotification {
	return &EmailNotification{Notification{&Email{}}}
}
