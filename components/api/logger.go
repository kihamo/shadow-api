package api

import (
	"sync"

	"github.com/kihamo/shadow/components/logger"
)

type Logger struct {
	l       logger.Logger
	mutext  sync.RWMutex
	enabled bool
}

func NewLogger(l logger.Logger) *Logger {
	return &Logger{
		l:       l,
		enabled: true,
	}
}

func (l *Logger) On() {
	l.mutext.Lock()
	defer l.mutext.Unlock()

	l.enabled = true
}

func (l *Logger) Off() {
	l.mutext.Lock()
	defer l.mutext.Unlock()

	l.enabled = false
}

func (l *Logger) Println(v ...interface{}) {
	l.mutext.RLock()
	defer l.mutext.RUnlock()

	if l.enabled {
		l.l.Info(v...)
	}
}

func (l *Logger) Printf(format string, v ...interface{}) {
	l.mutext.RLock()
	defer l.mutext.RUnlock()
	if l.enabled {
		l.l.Infof(format, v...)
	}
}
