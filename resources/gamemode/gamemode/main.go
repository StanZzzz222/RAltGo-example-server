package main

import (
	"fmt"
	"github.com/StanZzzz222/RAltGo/common/alt/alt_events"
	"github.com/StanZzzz222/RAltGo/common/alt/blip"
	"github.com/StanZzzz222/RAltGo/common/alt/ped"
	"github.com/StanZzzz222/RAltGo/common/alt/scheduler"
	"github.com/StanZzzz222/RAltGo/common/alt/timers"
	"github.com/StanZzzz222/RAltGo/common/alt/vehicle"
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/common/utils"
	"github.com/StanZzzz222/RAltGo/hash_enums"
	"github.com/StanZzzz222/RAltGo/hash_enums/ped_hash"
	"github.com/StanZzzz222/RAltGo/hash_enums/vehicle_hash"
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

/*
Benchmark
Basic API setting and acquisition performance testing
We found through rigorous testing that its performance is really good.
You can try the benchmark below. The following is the time taken by Intel Core i913900H
*/
func benchmark() {
	ped1 := ped.CreatePedByHash(ped_hash.Abner, utils.NewVector3(-1019.3187, -2928.9758, 14), utils.NewVector3(0, 0, 0))
	start := time.Now()
	for i := 0; i < 50000; i++ {
		_ = vehicle.CreateVehicleByHash(vehicle_hash.T20, fmt.Sprintf("test%v", i), utils.NewVector3(-1069.3187, -2928.9758, 14.1318), utils.NewVector3(0, 0, 0), 1, 1)
	}
	logger.LogInfof("Create 50,000 vehicles Since: %v ms", time.Since(start).Milliseconds()) // Since: 617 ms
	start = time.Now()
	for i := 0; i < 10000; i++ {
		_ = blip.CreateBlipPoint(12, 1, fmt.Sprintf("test%v", i), utils.NewVector3(-1069.3187, -2928.9758, 14.1318))
		_ = blip.CreateBlipRadius(13, 1, fmt.Sprintf("test%v", i), utils.NewVector3(-1069.3187, -2928.9758, 14.1318), 15)
		_ = blip.CreateBlipArea(14, 1, fmt.Sprintf("test%v", i), utils.NewVector3(-1069.3187, -2928.9758, 14.1318), 15, 15)
	}
	logger.LogInfof("Create 10,000 blips Since: %v ms", time.Since(start).Milliseconds()) // Since: 382 ms
	start = time.Now()
	for i := 0; i < 50000; i++ {
		_ = ped.CreatePedByHash(ped_hash.Agent, utils.NewVector3(-1069.3187, -2928.9758, 14), utils.NewVector3(0, 0, 0))
		_ = ped.CreatePedByHash(ped_hash.Ammucity01SMY, utils.NewVector3(-1089.3187, -2928.9758, 14), utils.NewVector3(0, 0, 0))
		_ = ped.CreatePedByHash(ped_hash.AcidLabCookIG, utils.NewVector3(-1039.3187, -2928.9758, 14), utils.NewVector3(0, 0, 0))
	}
	logger.LogInfof("Create 50,000 peds Since: %v ms", time.Since(start).Milliseconds()) // Since: 337 ms
	start = time.Now()
	for i := 0; i < 100000; i++ {
		_ = ped1.GetPosition()
		_ = ped1.GetRotation()
		ped1.SetMaxHealth(hash_enums.MaxHealth)
		ped1.SetPosition(utils.NewVector3(-1039.3187, -2928.9758, 14))
	}
	logger.LogInfof("Get position/rotation and set 100,000 peds Since: %v ms", time.Since(start).Milliseconds()) // Since: 781 ms
}

func onStart() {
	logger.LogInfo("Server start")
	// Benchmark
	benchmark()
}

func onStop() {
	logger.LogInfo("Server stop")
}

func onPlayerConnect(player *models.IPlayer) {
	logger.LogInfof("Player %v(%v) connected, IP: %v", player.GetName(), player.GetId(), player.GetIP())
	player.Spawn(ped_hash.FreemodeMale01, utils.NewVector3(-1069.3187, -2928.9758, 14.1318))
	timers.SetTimeout(time.Second*5, func() {
		player.SetPedModel(ped_hash.Ammucity01SMY)
		player.SetDateTimeUTC8(time.Now())
		logger.LogInfof("Change player %v model", player.GetName())
	})
	timers.SetTimeout(time.Second*8, func() {
		veh := vehicle.CreateVehicleByHash(vehicle_hash.T20, "RALTGO", utils.NewVector3(-1069.3187, -2928.9758, 14.1318), utils.NewVector3(0, 0, 0), 1, 1)
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
