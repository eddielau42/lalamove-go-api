package logger

import (
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	// 日志读写句柄
	logWriter *log.Logger
	// 日志输出路径
	output *os.File
	// 当前日志等级
	logLevel int
	// 当前日志文件名
	logFile string

	// 记录当前日期
	curDay int

	// 日志文件锁
	fileLock *sync.RWMutex
)

// 日志等级
const (
	DEBUG_LEVEL = iota
	INFO_LEVEL
	WARN_LEVEL
	ERROR_LEVEL
)

const (
	pathDepth = 3
	callerDepth = 4
)

func init() {
	fileLock = &sync.RWMutex{}
}

// SetLevel	设置日志等级
func SetLevel(level int) {
	logLevel = level
}

// SetFile	设置日志文件
func SetFile(file string) {
	var err error
	if output, err = os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664); err != nil {
		panic(err)
	} else {
		curDay = time.Now().YearDay()
		logFile = file
		logWriter = log.New(output, "", log.Ldate|log.Lmicroseconds)
	}
}

func Debug(format string, args ...any) {
	if logLevel <= DEBUG_LEVEL {
		WriteLog(DEBUG_LEVEL, format, args...)
	}
}
func Info(format string, args ...any) {
	if logLevel <= INFO_LEVEL {
		WriteLog(INFO_LEVEL, format, args...)
	}
}
func Warn(format string, args ...any) {
	if logLevel <= WARN_LEVEL {
		WriteLog(WARN_LEVEL, format, args...)
	}
}
func Error(format string, args ...any) {
	if logLevel <= ERROR_LEVEL {
		WriteLog(ERROR_LEVEL, format, args...)
	}
}

// WriteLog	日志内容写入
func WriteLog(level int, format string, args ...any) {
	checkDayChange()

	logWriter.Printf(getLogLevelTag(level) + " " + getPrefix() + format, args...)
}

// getCallTrace	获取调用栈
func getCallTrace() (string, int) {
	// 函数名, 文件名, 行号, 是否异常
	_, filename, line, ok := runtime.Caller(callerDepth)
	if ok {
		return filename, line
	}
	return "", 0
}

// getPrefix	返回前缀
func getPrefix() string {
	filename, line := getCallTrace()

	// 防止文件路径过长, 对文件路径进行截取
	path := strings.Split(filename, "/")
	if len(path) > pathDepth {
		filename = strings.Join(path[len(path)-pathDepth:], "/")
	}

	return filename + "(line:" + strconv.Itoa(line) + ") "
}

// getLogLevelTag	返回日志等级对应标签
func getLogLevelTag(level int) string {
	switch level {
	case DEBUG_LEVEL:
		return "[DEBUG]"
	case INFO_LEVEL:
		return "[INFO]"
	case WARN_LEVEL:
		return "[WARN]"
	case ERROR_LEVEL:
		return "[ERROR]"
	}
	return "[-]"
}

// checkDayChange	检查日期变更
func checkDayChange() {
	fileLock.Lock()
	defer fileLock.Unlock()

	// 日期无变更
	day := time.Now().YearDay()
	if day == curDay {
		return
	}

	// 关闭日志文件
	output.Close()
	// 修改日志文件名; 添加昨日日期作为日志文件后缀 (如: log -> log.20060102)
	sufferfix := time.Now().Add(-24 * time.Hour).Format("20060102")
	os.Rename(logFile, logFile + "." + sufferfix)

	// 重新设置日志文件
	SetFile(logFile)
}