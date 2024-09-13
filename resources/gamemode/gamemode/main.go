package main

import (
	"github.com/StanZzzz222/RAltGo/common/alt/alt_events"
	"github.com/StanZzzz222/RAltGo/common/alt/blip"
	"github.com/StanZzzz222/RAltGo/common/alt/scheduler"
	"github.com/StanZzzz222/RAltGo/common/alt/timers"
	"github.com/StanZzzz222/RAltGo/common/alt/vehicle"
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/common/utils"
	"github.com/StanZzzz222/RAltGo/enums/ped"
	vehicleModelHash "github.com/StanZzzz222/RAltGo/enums/vehicle"
	"github.com/StanZzzz222/RAltGo/logger"
	"github.com/StanZzzz222/RAltGo/modules"
	"time"
)

/*
   Create by zyx
   Date Time: 2024/9/10
   File: gamemode.go
*/

func main() {}
func init() {
	modules.InitMounted()
	alt_events.OnStart(onStart)
	alt_events.OnStop(onStop)
	alt_events.OnPlayerConnect(onPlayerConnect)
	alt_events.OnEnterVehicle(onEnterVehicle)
	alt_events.OnLeaveVehicle(onLeaveVehicle)
}

func onStart() {
	logger.LogInfo("Server start")
	b := blip.CreateBlipPoint(12, 1, "测试", utils.NewVector3(-1069.3187, -2928.9758, 14.1318))
	timers.SetTimeout(time.Second*3, func() {
		b.SetBlipFlashInterval(3)
		b.SetBlipFlashes(true)
	})
}

func onStop() {
	logger.LogInfo("Server stop")
}

func onPlayerConnect(player *models.IPlayer) {
	logger.LogInfof("Player %v(%v) connected, IP: %v", player.GetName(), player.GetId(), player.GetIP())
	player.Spawn(ped.FreemodeMale01, utils.NewVector3(-1069.3187, -2928.9758, 14.1318))
	timers.SetTimeout(time.Second*5, func() {
		player.SetPedModel(ped.Ammucity01SMY)
		player.SetDateTimeUTC8(time.Now())
		logger.LogInfof("Change player %v model", player.GetName())
	})
	timers.SetTimeout(time.Second*8, func() {
		veh := vehicle.CreateVehicleByHash(vehicleModelHash.T20, "RALTGO", utils.NewVector3(-1069.3187, -2928.9758, 14.1318), utils.NewVector3(0, 0, 0), 1, 1)
		logger.LogInfof("Create vehicle %v | model: %v", veh.GetId(), veh.GetModel())
		veh.SetNeonColor(utils.NewRGBA(123, 104, 238, 255))
		veh.SetNeonActive(false)
		player.SetIntoVehicle(veh, 1)
	})
}

func onEnterVehicle(player *models.IPlayer, vehicle *models.IVehicle, seat uint8) {
	logger.LogInfof("Player %v enter vehicle: %v", player.GetName(), vehicle.GetModel())
	vehicle.SetPrimaryColor(5)
	vehicle.SetNeonActive(true)
}

func onLeaveVehicle(player *models.IPlayer, vehicle *models.IVehicle, seat uint8) {
	logger.LogInfof("Player %v leave vehicle: %v", player.GetName(), vehicle.GetModel())
	s := scheduler.NewScheduler()
	s.AddTask(func() {
		vehicle.SetPrimaryColor(1)
		vehicle.SetNeonActive(false)
	})
	s.Run()
	logger.LogInfof("Continue running, task send to ontick scheduler")
}
