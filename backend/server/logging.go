package server

import "fmt"

/*
Levels:
0 - no logging
1 - server logs
2 - client and server logs
3 - debug logs
4 - error logs
*/
var logLevel int8 = 0

type LogMessage struct {
	Level   int8
	Message string
}

func SetLogLevel(ll int8) {
	logLevel = ll
}

func Log(logM LogMessage) {
	if logM.Level > logLevel {
		fmt.Println(logM.Message)
	}
}
