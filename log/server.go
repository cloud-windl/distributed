package log

import (
	stlog "log"
	"os"
)

var log *stlog.Logger

type fileLog string

func (fl fileLog) Write(data []byte) (n int, err error) {
	f, err := os.OpenFile(string(fl), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.Write(data)
}

func Run(destination string) {
	log = stlog.New(fileLog(destination), "[go]: ", stlog.LstdFlags)
}

func write(msg string) {
	log.Printf("%v\n", msg)
}
