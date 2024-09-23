package commands

import (
	"encoding/json"
	"fmt"
	"github.com/StanZzzz222/RAltGo/common/alt/command"
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/common/utils"
	"time"
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
		group.OnCommandDesc("emitbenchmark", EmitBenchmark, false, "/emitbenchmark [emit count] [user count]")
		group.OnCommandDesc("emitbenchmarkmap", EmitBenchmarkMaps, false, "/emitbenchmarkmap [emit count]")
		group.OnCommandDesc("getadmin", GetAdmin, false, "/getadmin [password]")
		group.OnCommandDesc("setpos", SetPos, false, "/setpos [Position (example: 0,0,0)]")
	}
}

func EmitBenchmark(player *models.IPlayer, emitCount, userCount int64) {
	type User struct {
		Id    int      `json:"id"`
		Name  string   `json:"name"`
		Age   int      `json:"age"`
		Likes []string `json:"likes"`
	}
	var users []*User
	for i := 0; i < int(userCount); i++ {
		users = append(users, &User{
			Id:    i + 1,
			Name:  fmt.Sprintf("User%d", i+1),
			Age:   i + 1,
			Likes: []string{"Like1", "Like2", "Like3"},
		})
	}
	userBytes, _ := json.Marshal(users)
	start := time.Now()
	for i := 0; i < int(emitCount); i++ {
		player.Emit("emitbenchmark", string(userBytes))
	}
	player.SendBroadcastMessage(fmt.Sprintf("Emit benchmark done, EmitCount: %v | UsersCount: %v | Since: %v ms", emitCount, userCount, time.Since(start).Milliseconds()))
}

func EmitBenchmarkMaps(player *models.IPlayer, emitCount int64) {
	start := time.Now()
	for i := 0; i < int(emitCount); i++ {
		player.Emit("emitbenchmark:objects", map[string]any{
			"id":    i + 1,
			"name":  fmt.Sprintf("User%d", i+1),
			"age":   i + 1,
			"likes": []string{"Like1", "Like2", "Like3"},
		})
	}
	player.SendBroadcastMessage(fmt.Sprintf("Emit benchmark done, EmitCount: %v  | Since: %v ms", emitCount, time.Since(start).Milliseconds()))
}

func Hello(player *models.IPlayer, name string, age int64) {
	player.Emit("hello", name, age)
}

func GetPos(player *models.IPlayer) {
	player.SendBroadcastMessage(fmt.Sprintf("Position: %v | Rotation: %v", player.GetPosition().ToString(), player.GetRotation().ToString()))
}

func GetAdmin(player *models.IPlayer, password string) {
	if password == "raltgo" {
		player.SetData("admin", true)
		player.SendBroadcastMessage("Take admin success")
	}
}

func SetPos(player *models.IPlayer, posStr string) {
	pos, err := utils.NewVector3FromStr(posStr)
	if err != nil {
		player.SendBroadcastMessage("Set position error, position incorrect format")
		return
	}
	player.SetPosition(pos)
}
