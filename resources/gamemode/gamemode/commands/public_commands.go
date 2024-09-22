package commands

import (
	"fmt"
	"github.com/StanZzzz222/RAltGo/common/alt/command"
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/common/utils"
)

/*
   Create by zyx
   Date Time: 2024/9/22
   File: public_commands.go
*/

func InitPublicCommands() {
	group := command.NewCommandGroup("PublicCommands")
	// Add OnCommand
	{
		group.OnCommand("getpos", GetPos, false)
		group.OnCommandDesc("hello", Hello, false, "/hello [name] [age]")
		group.OnCommandDesc("getadmin", GetAdmin, false, "/getadmin [password]")
		group.OnCommandDesc("setpos", SetPos, false, "/setpos [Position (example: 0,0,0)]")
	}
}

func Hello(player *models.IPlayer, name string, age int64) {
	player.Emit("hello", name, age)
}

func GetPos(player *models.IPlayer) {
	player.SendBroadcast(fmt.Sprintf("Position: %v | Rotation: %v", player.GetPosition().ToString(), player.GetRotation().ToString()))
}

func GetAdmin(player *models.IPlayer, password string) {
	if password == "raltgo" {
		player.SetData("admin", true)
		player.SendBroadcast("Take admin success")
	}
}

func SetPos(player *models.IPlayer, posStr string) {
	pos, err := utils.NewVector3FromStr(posStr)
	if err != nil {
		player.SendBroadcast("Set position error, position incorrect format")
		return
	}
	player.SetPosition(pos)
}
