package common

import (
	"os"
	"os/signal"
	"syscall"
)

// WaitStopSignal 阻塞 stop 系统信号
func WaitStopSignal() {
	WaitSignals(os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT)
}

// WaitSignals 阻塞等待系统信号
func WaitSignals(signals ...os.Signal) {
	signalChan := make(chan os.Signal)
	defer func() {
		signal.Stop(signalChan)
		close(signalChan)
	}()
	signal.Notify(signalChan, signals...)
	<-signalChan
}
