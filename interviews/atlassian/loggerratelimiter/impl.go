package loggerratelimiter

type Logger struct {
	msgMap map[string]int
}

func Constructor() Logger {
	return Logger{
		msgMap: make(map[string]int),
	}
}

func (l *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	if ts, ok := l.msgMap[message]; !ok || timestamp >= ts {
		l.msgMap[message] = timestamp + 10
		return true
	}
	return false
}
