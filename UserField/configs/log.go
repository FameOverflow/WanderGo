package Config

import (
	"log"
	"os"
)

//生成日志
func InitLogging() {
	logFile, _ := os.OpenFile("RinaLog.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	log.SetOutput(logFile)
	log.SetPrefix("[Debug] -")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
