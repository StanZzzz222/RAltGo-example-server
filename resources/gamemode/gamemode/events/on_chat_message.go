package events

import (
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/logger"
)

/*
   Create by zyx
   Date Time: 2024/9/22
   File: on_chat_message.go
*/

func OnChatMessage(player *models.IPlayer, message string) {
	logger.LogInfof("Player: %v say %v", player.GetName(), message)
}
