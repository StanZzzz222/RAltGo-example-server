package commands

import (
	"encoding/json"
	"fmt"
	"github.com/StanZzzz222/RAltGo/common"
	"github.com/StanZzzz222/RAltGo/common/alt/timers"
	"github.com/StanZzzz222/RAltGo/common/alt/vehicle"
	"github.com/StanZzzz222/RAltGo/common/command"
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/hash_enums/radio_station_type"
	"github.com/StanZzzz222/RAltGo/hash_enums/vehicle_hash"
	"github.com/StanZzzz222/RAltGo/hash_enums/vehicle_light_state_type"
	"github.com/StanZzzz222/RAltGo/hash_enums/vehicle_lock_state_type"
	"math/rand"
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
		group.OnCommandDesc("sayhi", SayHi, true, "/sayhi [content]")
		group.OnCommandDesc("emitbenchmark", EmitBenchmark, false, "/emitbenchmark [emit count] [user count]")
		group.OnCommandDesc("emitbenchmarkmap", EmitBenchmarkMaps, false, "/emitbenchmarkmap [emit count]")
		group.OnCommandDesc("basebenchmark", BaseBenchmark, false, "/basebenchmark [type] [count]")
		group.OnCommandDesc("getadmin", GetAdmin, false, "/getadmin [password]")
		group.OnCommandDesc("setpos", SetPos, false, "/setpos [Position (example: 0,0,0)]")
	}
}

func BaseBenchmark(player *models.IPlayer, t int64, count int64) {
	switch t {
	case 0:
		// Mock real usage environment
		// Intel Core i913900H - count: 10000 - Since: 135 ms | Game - Since: 800 ms
		pos := common.NewVector3ARound(player.GetPosition().X, player.GetPosition().Y, player.GetPosition().Z, 10)
		veh := vehicle.CreateVehicleByHash(vehicle_hash.T20, "test", pos, player.GetRotation(), 1, 1)
		player.SetIntoVehicle(veh, 1)
		timers.SetTimeout(time.Second*3, func() {
			start := time.Now()
			for i := 0; i < int(count); i++ {
				veh.SetPrimaryColor(uint8(rand.Intn(159)))
				veh.SetSecondColor(uint8(rand.Intn(159)))
				if i%2 == 0 {
					veh.SetEngineOn(true)
					veh.SetNeonActive(true)
					veh.SetLockState(vehicle_lock_state_type.VehicleLockLocked)
					veh.SetDriftMode(true)
				} else {
					veh.SetEngineOn(false)
					veh.SetNeonActive(false)
					veh.SetLockState(vehicle_lock_state_type.VehicleLockUnlocked)
					veh.SetDriftMode(true)
				}
				veh.Repair()
				veh.SetWheelColor(uint8(rand.Intn(10)))
				veh.SetNeonColor(common.NewRGBA(uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255))
				veh.SetLightState(vehicle_light_state_type.VehicleLightAlwaysOn)
				veh.SetDirtLevel(uint8(rand.Intn(6)))
				veh.SetRadioStation(radio_station_type.LosSantosRockRadio)
				veh.SetDashboardColor(uint8(rand.Intn(6)))
				_ = veh.GetPosition()
				_ = veh.GetRotation()
				veh.ToggleEngine()
			}
			player.SendBroadcastMessage(fmt.Sprintf("Vehicle benchmark, Count: %v | Since: %v ms", count, time.Since(start).Milliseconds()))
		})
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

func SayHi(player *models.IPlayer, content string) {
	player.SendBroadcastMessage(fmt.Sprintf("* Hi %s", content))
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
	pos, err := common.NewVector3FromStr(posStr)
	if err != nil {
		player.SendBroadcastMessage("Set position error, position incorrect format")
		return
	}
	player.SetPosition(pos)
}
