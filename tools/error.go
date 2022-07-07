package tools

import (
	"log"
	"runtime"
)

func Log(v ...interface{}) {
	caller := 1
	for {
		if _, file, line, ok := runtime.Caller(caller); ok {
			log.Println("warning...", v, file, line)
			caller++
		} else {
			return
		}
	}
}
