package events

import (
	"github.com/StanZzzz222/RAltGo/common"
	"github.com/StanZzzz222/RAltGo/common/core/alt/alt_timers"
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/hash_enums/ped_hash"
	"github.com/StanZzzz222/RAltGo/logger"
	"time"
)

/*
   Create by zyx
   Date Time: 2024/9/22
   File: on_player_connect.go
*/

func OnPlayerConnect(player *models.IPlayer) {
	logger.Logger().LogInfof("Player %v(%v) connected, IP: %v", player.GetName(), player.GetId(), player.GetIP())
	player.SendBroadcastMessage("Welcome to server")
	player.Spawn(ped_hash.FreemodeMale01, common.NewVector3(-1069.3187, -2928.9758, 14.1318))
	alt_timers.SetTimeout(time.Second*3, func() {
		player.SetPedModel(ped_hash.AnitaCutscene)
		player.SetSyncedMetaData("test", "Hello JS")
		player.SetSyncedMetaData("bool", true)
		player.SetSyncedMetaData("player", player)
		player.Emit("test_meta_data")
	})
}
