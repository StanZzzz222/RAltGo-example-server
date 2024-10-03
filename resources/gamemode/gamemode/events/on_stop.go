package events

import "github.com/StanZzzz222/RAltGo/logger"

/*
   Create by zyx
   Date Time: 2024/9/22
   File: on_start.go
*/

func OnStop() {
	logger.Logger().LogInfo("Server Stop")
}
