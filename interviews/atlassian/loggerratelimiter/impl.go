package loggerratelimiter

import "sync"

type Logger struct {
	mutex  sync.Mutex
	msgMap map[string]int
}

func Constructor() Logger {
	return Logger{
		mutex:  sync.Mutex{},
		msgMap: make(map[string]int),
	}
}

func (l *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	l.mutex.Lock()
	defer l.mutex.Unlock() // Ensure mutex is unlocked after the operation

	// As msgMap map[string]int, the ts has the default value 0
	if ts, ok := l.msgMap[message]; !ok || timestamp >= ts {
		l.msgMap[message] = timestamp + 10
		return true
	}
	return false
}

func (l *Logger) ShouldPrintMessage2(timestamp int, message string) bool {
	if allowedTime, found := l.msgMap[message]; found {
		if timestamp >= allowedTime {
			l.msgMap[message] = timestamp + 10
			return true
		}
		return false
	}
	l.msgMap[message] = timestamp + 10
	return true
}
