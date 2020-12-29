package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 往文件里写入日志
// FileLogger 文件日志结构体

//根据指定的日志文件路径和文件名打开日志文件
func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed ,err:%s", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed ,err:%s", err)
		return err
	}
	//日志文件都已经打开了
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	//开启5个后台goroutine往文件里写日志：会出现多个goroutine占用一个文件的情况
	//for i:=0;i<5;i++{
	//	go f.writeLogBackground()
	//}
	//开启1个后台goroutine往文件里写日志
	go f.writeLogBackground()
	return nil
}

//关闭文件
func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}

var (
	// 日志通道的大小
	MaxChanSize = 50000
)

// NewFileLogger构造函数
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
		logChan:     make(chan *LogMsg, MaxChanSize),
	}
	err = fl.initFile()
	if err != nil {
		panic(err)
	}
	return fl
}

//判断是否需要记录日志
func (f *FileLogger) enable(loglevel LogLevel) bool {
	return loglevel >= f.Level
}

// 判断文件是否需要切割
func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("get file failed,err:", err)
		return false
	}
	//如果当前文件大小 大于等于 日志文件的最大值，返回true
	return fileInfo.Size() >= f.maxFileSize
}

// 切割文件
func (f *FileLogger) spiltFile(file *os.File) (*os.File, error) {
	//需要切割日志文件
	nowStr := time.Now().Format("20060102150405000")
	//文件名称取传的参数file
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("get file failed,err:", err)
		return nil, err
	}
	//无论是log文件还是log.err文件都能拿到对应名字：fileInfo.Name()
	logName := path.Join(f.filePath, fileInfo.Name())
	//logName := path.Join(f.filePath,f.fileName)  //这个写法拿到的是文件本身的name
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr)
	//1、关闭当前日志文件
	file.Close()
	//2、备份一下 rename
	os.Rename(logName, newLogName)
	//3、打开一个新的日志文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("opne new log file failed ,err:%s\n", err)
		return nil, err
	}

	//4、将打开的新文件对象赋值给 f.fileObj
	return fileObj, nil
}

func (f *FileLogger) writeLogBackground() {
	for {
		if f.checkSize(f.fileObj) {
			newFile, err := f.spiltFile(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
		select {
		case logTmp := <-f.logChan:
			logInfo := fmt.Sprintf("[%s] [%s] [%s:%s:line%d] %s \n", logTmp.Timestamp, getLogString(logTmp.Level), logTmp.FileName, logTmp.funcName, logTmp.Line, logTmp.Msg)
			fmt.Fprintf(f.fileObj, logInfo)
			if logTmp.Level >= ERROR {
				if f.checkSize(f.errFileObj) {
					newFile, err := f.spiltFile(f.errFileObj)
					if err != nil {
						return
					}
					f.errFileObj = newFile
				}
				// 如果记录的日志大于等于error级别，还要再error日志文件中再记录一遍
				fmt.Fprintf(f.errFileObj, logInfo)
			}
		default:
			//取不到日志则睡500ms
			time.Sleep(time.Millisecond * 500)
		}
	}
}

//记录日志
func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getINFO(3)

		//先把日志放到通道中
		LogTmp := &LogMsg{
			Level:     lv,
			Msg:       msg,
			funcName:  funcName,
			FileName:  fileName,
			Timestamp: now.Format("2006-01-02 15:04:05"),
			Line:      lineNo,
		}
		select {
		case f.logChan <- LogTmp:
		default:
			//满了日志丢掉，保证日志不出现阻塞
		}

	}

}

//DEBUG方法
func (f *FileLogger) DEBUG(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}

func (f *FileLogger) TRACE(format string, a ...interface{}) {
	f.log(TRACE, format, a...)
}

func (f *FileLogger) INFO(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}

func (f *FileLogger) WARNING(format string, a ...interface{}) {
	f.log(WARNING, format, a...)
}

func (f *FileLogger) ERROR(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}

func (f *FileLogger) FATAL(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}
