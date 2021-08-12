// Logger это набор для упрощенного подключения логирования к проекту или пакету
// Его использование ведет к уменьшению повторения boilerplate кода в проектах

// Пакет в своей основе использует Zap - невероятно быстрый, структурированный,
// с поддержкой уровней логгер, разработанный в Uber: https://github.com/uber-go/zap

package logger

import (
	"log"
	"strings"

	"go.uber.org/zap"
)

// AtomicLevel уровень вывода логов
var AtomicLevel zap.AtomicLevel

// l внутренняя переменная для хранения глобального объекта текущего логгера
// по умолчанию используется SugaredLogger
var l *Log

func init() {
	config := zap.NewProductionConfig()
	AtomicLevel = config.Level

	// TODO Написать более конфигурируемую инициализацию из сервиса или модуля (используя Options, кастомный вывод логов)
	logger, err := config.Build()
	if err != nil {
		log.Fatalf("Could not initialize a Zap logger: %v\n", err)
	}

	zap.ReplaceGlobals(logger)

	l = New()
}

// New создает новый объект нашего логгера
func New() *Log {
	return &Log{Logger: zap.S()}
}

// Log структура логгера с использованным по умолчанию Sugar версией логгера Zap
type Log struct {
	Logger *zap.SugaredLogger
}

// Instance return Sugar instance of a Zap logger
func Instance() *zap.SugaredLogger { return l.Instance() }

// Instance return Sugar instance of a Zap logger
func (l *Log) Instance() *zap.SugaredLogger {
	return l.Logger
}

// Plain returns plain instance of a Zap logger
func Plain() *zap.Logger { return l.Plain() }

// Plain returns plain instance of a Zap logger
func (l *Log) Plain() *zap.Logger {
	return l.Logger.Desugar()
}

// SetLevel хелпер для установки уровня фильтрации логов
// Подробнее про уровни: https://www.godoc.org/go.uber.org/zap#pkg-constants
func SetLevel(level string) { l.SetLevel(level) }

// SetLevel хелпер для установки уровня фильтрации логов
func (l *Log) SetLevel(level string) {
	level = strings.ToLower(level)
	switch level {
	case "debug":
		AtomicLevel.SetLevel(zap.DebugLevel)
	case "info":
		AtomicLevel.SetLevel(zap.InfoLevel)
	case "warn":
		AtomicLevel.SetLevel(zap.WarnLevel)
	case "error":
		AtomicLevel.SetLevel(zap.ErrorLevel)
	case "dpanic":
		AtomicLevel.SetLevel(zap.DPanicLevel)
	case "panic":
		AtomicLevel.SetLevel(zap.PanicLevel)
	case "fatal":
		AtomicLevel.SetLevel(zap.FatalLevel)
	default:
		AtomicLevel.SetLevel(zap.InfoLevel)
		log.Printf("Couldn't find a proper zap log level for %s.\nSetting up a resonable default: InfoLevel\n", level)
	}
}

// Debug хелпер для установки в приложении (сервисе) уровня логов для дебаггинга
func Debug(level string) { l.Debug(level) }

// Debug хелпер для установки в приложении (сервисе) уровня логов для дебаггинга
func (l *Log) Debug(level string) {
	config := zap.NewDevelopmentConfig()
	AtomicLevel = config.Level

	logger, err := config.Build()
	if err != nil {
		log.Fatalf("Could not initalize a Zap logger instance: %v\n", err)
	}

	if level == "" {
		level = "debug"
	}
	l.SetLevel(level)
	l.Logger = logger.Sugar()

	zap.ReplaceGlobals(logger)
}

// Flush performs a core sync method to clear all buffered log entries
func Flush() { l.Flush() }

// Flush performs a core sync method to clear all buffered log entries
func (l *Log) Flush() {
	_ = l.Logger.Sync()
}
