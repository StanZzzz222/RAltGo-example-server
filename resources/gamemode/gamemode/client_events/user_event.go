package client_events

import (
	"fmt"
	"github.com/StanZzzz222/RAltGo/common/alt/alt_events"
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/logger"
)

/*
   Create by zyx
   Date Time: 2024/9/22
   File: user_event.go
*/

func InitUserEvents() {
	alt_events.Events().OnClientEvent("hello", HelloEvent)
}

func HelloEvent(player *models.IPlayer, name string, age int64) {
	logger.LogInfof("Trigger HelloEvent Hello %v, age is %v", name, age)
	player.SendBroadcastMessage(fmt.Sprintf("Hello %v, age is %v", name, age))
}
