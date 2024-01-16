package logging

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

var (
	logger     *logrus.Entry
	loggerFile *os.File
)

func GetLogger() (*logrus.Entry, *os.File) {
	return logger, loggerFile
}

func InitLogger() {
	log := logrus.New()                    // создание экземпляра объекта logrus
	log.SetReportCaller(true)              // записывать в лог информацию о пользователе, отправившим запрос
	log.SetLevel(logrus.TraceLevel)        // лог для информарования
	log.Formatter = &logrus.TextFormatter{ // изменение форматирования логгера
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s:%d", filename, f.Line), fmt.Sprintf("%s()", f.Function)
		},
		DisableColors: false,
		FullTimestamp: true,
	}

	err := os.MkdirAll("logs", 0660)
	if err != nil {
		panic("[Error]: can't make `logs` dir")
	} else {
		loggerFile, err = os.OpenFile("logs/log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
		if err != nil {
			panic(fmt.Sprintf("[Error]: %s", err))
		}

		log.SetOutput(loggerFile) // устанавливаем файл для лога
	}

	logger = logrus.NewEntry(log)

}
