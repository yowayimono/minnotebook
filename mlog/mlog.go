package mlog

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

type LogLevel int

const (
	Info LogLevel = iota
	Debug
	Warning
	Error
)

var logLevelColors = map[LogLevel]string{
	Debug:   "\033[0;37m", // Gray
	Info:    "\033[0;32m", // Green
	Warning: "\033[0;33m", // Yellow
	Error:   "\033[0;31m", // Red
}

const ColorReset = "\033[0m"

type Logger struct {
	level  LogLevel
	ToFile bool
	file   *os.File

	asyncLogs   chan string
	filename    string
	entry       LogEntry
	nextLogTime time.Time
	jsonFormat  bool // 是否使用 JSON 格式
	mutex       sync.Mutex
}

type LogEntry struct {
	data map[string]interface{}
}

// NewJsonLogger 用于创建一个新的 JSON 格式的 Logger 实例
func NewJsonLogger(level LogLevel, filename string) (*Logger, error) {
	logger := &Logger{
		level:       level,
		ToFile:      filename != "",
		asyncLogs:   make(chan string, 10000),
		filename:    filename,
		entry:       LogEntry{data: make(map[string]interface{})},
		nextLogTime: time.Now().Truncate(24 * time.Hour).Add(24 * time.Hour),
		jsonFormat:  true, // 设置为使用 JSON 格式
	}

	if logger.ToFile {
		if err := logger.openLogFile(); err != nil {
			return nil, err
		}
		go logger.startAsyncLogWriter()
	}

	return logger, nil
}

func NewLogger(level LogLevel, filename string) (*Logger, error) {
	logger := &Logger{
		level:       level,
		ToFile:      filename != "",
		asyncLogs:   make(chan string, 10000),
		filename:    filename,
		nextLogTime: time.Now().Truncate(24 * time.Hour).Add(24 * time.Hour),
	}

	if logger.ToFile {
		if err := logger.openLogFile(); err != nil {
			return nil, err
		}
		go logger.startAsyncLogWriter()
	}

	return logger, nil
}

func (l *Logger) startAsyncLogWriter() {
	for logEntry := range l.asyncLogs {
		l.mutex.Lock()
		if l.file != nil {
			if time.Now().After(l.nextLogTime) {
				fmt.Println("进来啦1111")
				l.rotateLogFile()
			}
			l.file.WriteString(logEntry)
		}
		l.mutex.Unlock()
	}
}

func (l *Logger) rotateLogFile() {

	l.file.Close()

	l.filename = filepath.Join(filepath.Dir(l.filename), time.Now().Format("2006-01-02-13:02")+".log")
	l.openLogFile()

	l.nextLogTime = l.nextLogTime.Add(24 * time.Hour)

}

func (l *Logger) openLogFile() error {
	file, err := os.OpenFile(l.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	l.file = file
	return nil
}

func (l *Logger) log(level LogLevel, message string) {
	if level < l.level {
		return
	}

	//l.mutex.Lock()
	//defer l.mutex.Unlock()

	now := time.Now()

	if !l.ToFile {
		logLine := fmt.Sprintf("[MIN] |%s%s%s| %s  %s       ——> %s\n", getColorCode(level), levelToString(level), ColorReset, now.Format("2006/01/02-15:04:05"), getCallerInfo(), message)
		fmt.Printf(logLine)
		return
	}
	if l.jsonFormat {
		logEntry := map[string]interface{}{
			"level":     levelToString(level),
			"timestamp": now.Format(time.RFC3339),
			"caller":    getCallerInfo(),
			"message":   message,
		}

		jsonData, err := json.Marshal(logEntry)
		if err != nil {
			fmt.Printf("Error encoding log entry to JSON: %v\n", err)
			return
		}

		l.asyncLogs <- string(jsonData) + "\n"
		return
	}

	if l.file != nil {
		logLine := fmt.Sprintf("[%s] [%s] [%s] ——> %s\n", levelToString(level), now.Format("2006-01-02 15:04:05"), getCallerInfo(), message)
		l.asyncLogs <- logLine
		l.file.Sync()
	}
}

func (l *Logger) Num(key string, value interface{}) *Logger {
	l.entry.data[key] = value
	return l
}

func (l *Logger) Str(key, value string) *Logger {
	l.entry.data[key] = value
	return l
}
func (l *Logger) Message(message string) {
	if l.level < l.level {
		return
	}

	logEntry := map[string]interface{}{
		"level":     levelToString(l.level),
		"timestamp": time.Now().Format(time.RFC3339),
		"message":   message,
		"data":      l.entry.data,
	}

	jsonData, err := json.Marshal(logEntry)
	if err != nil {
		fmt.Printf("Error encoding log entry to JSON: %v\n", err)
		return
	}

	fmt.Println(string(jsonData))
}

// Debug 写入 Debug 级别日志
func (l *Logger) Debug(message string) {
	l.log(Debug, message)

}

// Info 写入 Info 级别日志
func (l *Logger) Info(message string) {
	l.log(Info, message)
}

// Warning 写入 Warning 级别日志
func (l *Logger) Warning(message string) {
	l.log(Warning, message)
}

// Error 写入 Error 级别日志
func (l *Logger) Error(message string) {
	l.log(Error, message)
}

// Close 关闭日志文件
func (l *Logger) Close() error {
	if l.file != nil {
		close(l.asyncLogs) // 关闭异步通道，通知异步协程停止
		return l.file.Close()
	}
	return nil
}

func Log(level LogLevel, message string) error {
	logger := NewConsoleLogger(level)
	logger.log(level, message)
	return nil
}

// levelToString 将日志等级转换为字符串
func levelToString(level LogLevel) string {
	switch level {
	case Debug:
		return "DEBUG"
	case Info:
		return "INFO"
	case Warning:
		return "WARNING"
	case Error:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

// getColorCode 获取日志等级对应的颜色码
func getColorCode(level LogLevel) string {
	if color, ok := logLevelColors[level]; ok {
		return color
	}
	return ""
}

// NewConsoleLogger 用于创建一个新的控制台 Logger 实例
func NewConsoleLogger(level LogLevel) *Logger {
	return &Logger{
		level:  level,
		ToFile: false,
	}
}

// getCallerInfo 获取调用者信息
func getCallerInfo() string {
	_, file, line, ok := runtime.Caller(3) // Adjust the depth to get the caller correctly
	if !ok {
		file = "unknown"
		line = 0
	}
	caller := fmt.Sprintf("%s:%d", path.Base(file), line)
	return caller
}
