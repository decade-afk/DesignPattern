package structuralpatterns

import "fmt"

type Logger interface {
	Log(message string)
}

type ThirdPartyLogger interface {
	LogMessage(msg string)
}

type MyThirdPartyLogger struct {
}

func (l *MyThirdPartyLogger) LogMessage(msg string) {
	// Do something with the message
	fmt.Println("ThirdPartyLogger: " + msg)
}

type ThirdPartyLoggerAdapter struct {
	Logger ThirdPartyLogger
}

func (a *ThirdPartyLoggerAdapter) Log(message string) {
	a.Logger.LogMessage(message)
}
