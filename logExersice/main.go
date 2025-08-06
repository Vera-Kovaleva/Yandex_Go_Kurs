package main

import "log"

type LogLevel int

const (
	LogLevelError LogLevel = iota
	LogLevelWarning
	LogLevelInfo
)

type LogExtended struct {
	*log.Logger
	logLevel LogLevel
}

func NewLogExtended(level LogLevel) LogExtended {
	return LogExtended{
		Logger:   log.Default(),
		logLevel: level,
	}
}

func (l LogLevel) IsValid() bool {
	switch l {
	case LogLevelInfo, LogLevelWarning, LogLevelError:
		return true
	default:
		return false
	}
}

func (l *LogExtended) SetLogLevel(logLvl LogLevel) {
	if !logLvl.IsValid() {
		return
	}
	l.logLevel = logLvl
}

func (l *LogExtended) println(srcLogLvl LogLevel, prefix, msg string) {
	if l.logLevel < srcLogLvl {
		return
	}
	l.Logger.Println(prefix + msg)
}

func (l LogExtended) Infoln(msg string) {
	l.println(LogLevelInfo, "INFO ", msg)
}

func (l LogExtended) Warnln(msg string) {
	l.println(LogLevelWarning, "WARNING ", msg)

}

func (l LogExtended) Errorln(msg string) {
	l.println(LogLevelError, "ERROR ", msg)

}

func main() {
	logger := NewLogExtended(1)
	logger.SetLogLevel(LogLevelWarning)
	logger.Infoln("Не должно напечататься")
	logger.Warnln("Hello")
	logger.Errorln("World")
	logger.Println("Debug")
}
