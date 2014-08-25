// Package: fileLogger
// File: writer.go
// Created by: mint(mint.zhao.chiu@gmail.com)_aiwuTech
// Useage: 
// DATE: 14-8-24 12:40
package fileLogger

import (
	"log"
	"fmt"
	"time"
	"runtime"
)

const (
	DEFAULT_PRINT_INTERVAL = 300
)

// Receive logStr from f's logChan and print logstr to file
func (f *FileLogger) logWriter() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("FileLogger's LogWritter() catch panic: %v\n", err)
		}
	}()

	//TODO let printInterVal can be configure, without change sourceCode
	printInterval := DEFAULT_PRINT_INTERVAL

	seqTimer := time.NewTicker(time.Duration(printInterval) * time.Second)
	for {
		select {
		case str := <- f.logChan:

			f.p(str)
		case <- seqTimer.C:
			log.Printf("================ LOG SEQ SIZE:%v ==================\n", len(f.logChan))
		}
	}
}

// print log
func (f *FileLogger) p(str string) {
	f.mu.RLock()
	defer f.mu.RUnlock()

	f.lg.Output(2, str)
}

// Printf throw logstr to channel to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
func (f *FileLogger) Printf(format string, v ...interface {}) {
	f.logChan <- fmt.Sprintf(format, v...)
}

// Print throw logstr to channel to print to the logger.
// Arguments are handled in the manner of fmt.Print.
func (f *FileLogger) Print(v ...interface {}) {
	f.logChan <- fmt.Sprint(v...)
}

// Println throw logstr to channel to print to the logger.
// Arguments are handled in the manner of fmt.Println.
func (f *FileLogger) Println(v ...interface {}) {
	f.logChan <- fmt.Sprintln(v...)
}

//======================================================================================================================
// Trace log
func (f *FileLogger) Trace(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(2) //calldepth=3
	if f.logLevel <= TRACE {
		f.logChan <- fmt.Sprintf("[%v:%v]", shortFileName(file), line) + fmt.Sprintf("\033[1;35m[TRACE] "+format+" \033[0m ", v...)
	}
}

// same with Trace()
func (f *FileLogger) T(format string, v ...interface{}) {
	f.Trace(format, v...)
}

// info log
func (f *FileLogger) Info(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1) //calldepth=3
	if f.logLevel <= INFO {
		f.logChan <- fmt.Sprintf("[%v:%v]", shortFileName(file), line) + fmt.Sprintf("[INFO] "+format, v...)
	}
}

// same with Info()
func (f *FileLogger) I(format string, v ...interface{}) {
	f.Info(format, v...)
}

// warning log
func (f *FileLogger) Warn(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1) //calldepth=3
	if f.logLevel <= WARN {
		f.logChan <- fmt.Sprintf("[%v:%v]", shortFileName(file), line) + fmt.Sprintf("\033[1;33m[WARN] "+format+" \033[0m ", v...)
	}
}

// same with Warn()
func (f *FileLogger)W(format string, v ...interface{}) {
	f.Warn(format, v...)
}

// error log
func (f *FileLogger) Error(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1) //calldepth=3
	if f.logLevel <= ERROR {
		f.logChan <- fmt.Sprintf("%v:%v]", shortFileName(file), line) + fmt.Sprintf("\033[1;4;31m[ERROR] "+format+" \033[0m ", v...)
	}
}

// same with Error()
func (f *FileLogger) E(format string, v ...interface{}) {
	f.Error(format, v...)
}