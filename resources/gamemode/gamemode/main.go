package main

import (
	"fmt"
	"gamemode/client_events"
	"gamemode/commands"
	"gamemode/events"
	"github.com/StanZzzz222/RAltGo/common/alt/alt_events"
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/modules"
)

/*
   Create by zyx
   Date Time: 2024/9/10
   File: gamemode.go
*/

func main() {}
func init() {
	// Init go modules
	// Note: Do not delete! Do not delete! Do not delete!
	modules.InitMounted()
	// Base server events
	alt_events.Events().OnStart(events.OnStart)
	alt_events.Events().OnStop(events.OnStop)
	alt_events.Events().OnPlayerConnect(events.OnPlayerConnect)
	alt_events.Events().OnEnterVehicle(events.OnPlayerEnterVehicle)
	alt_events.Events().OnLeaveVehicle(events.OnLeaveVehicle)
	alt_events.Events().OnChatMessage(events.OnChatMessage)
	alt_events.Events().OnNetOwnerChange(func(entity any, oldNetOwner *models.IPlayer, newNetOwner *models.IPlayer) {
		fmt.Println(entity)
		fmt.Println(oldNetOwner)
		fmt.Println(newNetOwner)
	})
	// Client events
	client_events.InitUserEvents()
	// Commands
	commands.InitPublicCommands()
	commands.InitAdminCommands()
}
