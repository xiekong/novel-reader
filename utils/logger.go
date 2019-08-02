/**
 * @author XieKong
 * @date   2019/7/31 11:42
 */
package utils

import (
	"errors"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path/filepath"
	"time"
)

var Logger *logrus.Logger

func Init(debug bool) {
	if debug {
		initConsoleLogger()
	} else {
		initFileLogger()
	}
}

type logFileWriter struct {
	file *os.File
	size int64
}

func (writer *logFileWriter) Write(d []byte) (n int, err error) {
	if writer == nil {
		return 0, errors.New("logFileWriter is nil")
	}
	if writer.file == nil {
		return 0, errors.New("file not opened")
	}
	n, e := writer.file.Write(d)
	writer.size += int64(n)
	//文件最大 64K byte
	if writer.size > 1024*64 {
		writer.file.Close()
		writer.file, _ = os.OpenFile(filepath.Join(GetLogPath(), "novel-reader-"+time.Now().Format("20060102")+".log"), os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0600)
		writer.size = 0
	}
	return n, e
}

func initConsoleLogger() {
	Logger = logrus.New();
	Logger.SetFormatter(&logrus.TextFormatter{})
	Logger.SetLevel(logrus.DebugLevel)
	Logger.SetOutput(os.Stdout)
}

func initFileLogger() {
	file, err := os.OpenFile(filepath.Join(GetLogPath(), "novel-reader-"+time.Now().Format("20060102")+".log"), os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0600)
	if err != nil {
		log.Fatal("log file init failed")
	}

	info, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	Logger = logrus.New();
	Logger.SetFormatter(&logrus.TextFormatter{})
	Logger.SetLevel(logrus.DebugLevel)
	fileWriter := logFileWriter{file, info.Size()}
	Logger.SetOutput(&fileWriter)
}
