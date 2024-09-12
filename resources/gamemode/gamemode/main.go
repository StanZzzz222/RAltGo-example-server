package main

import (
	"github.com/StanZzzz222/RAltGo/alt_events"
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/common/utils"
	"github.com/StanZzzz222/RAltGo/enums/ped"
	vehicleModelHash "github.com/StanZzzz222/RAltGo/enums/vehicle"
	"github.com/StanZzzz222/RAltGo/logger"
	"github.com/StanZzzz222/RAltGo/modules"
	"github.com/StanZzzz222/RAltGo/scheduler"
	"github.com/StanZzzz222/RAltGo/timers"
	"github.com/StanZzzz222/RAltGo/vehicle"
	"sync"
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
	alt_events.OnLeaveVehicle(func(player *models.IPlayer, vehicle *models.IVehicle, seat uint8) {
		logger.LogInfof("Player %v leave vehicle: %v", player.GetName(), vehicle.GetModel())
		s := scheduler.NewScheduler()
		wg := &sync.WaitGroup{}
		go func() {
			wg.Add(1)
			s.AddTask(func() {
				vehicle.SetPrimaryColor(1)
				wg.Done()
			})
		}()
		wg.Wait()
		s.RunWait()
		logger.LogInfof("Done")
	})
}

func onStart() {
	logger.LogInfo("Server start")
}

func onStop() {
	logger.LogInfo("Server stop")
}

func onPlayerConnect(player *models.IPlayer) {
	logger.LogInfof("Player %v(%v) connected, IP: %v", player.GetName(), player.GetId(), player.GetIP())
	player.Spawn("mp_m_freemode_01", utils.NewVector3(-1069.3187, -2928.9758, 14.1318))
	timers.SetTimeout(time.Second*5, func() {
		player.SetPedModelByHash(ped.Ammucity01SMY)
		player.SetDateTimeUTC8(time.Now())
		logger.LogInfof("Change player %v model", player.GetName())
	})
	timers.SetTimeout(time.Second*8, func() {
		veh := vehicle.CreateVehicleByHash(vehicleModelHash.T20, "RALTGO", utils.NewVector3(-1069.3187, -2928.9758, 14.1318), utils.NewVector3(0, 0, 0), 1, 1)
		logger.LogInfof("Create vehicle %v | model: %v", veh.GetId(), veh.GetModel())
		veh.SetHeadLightColor(8)
		player.SetIntoVehicle(veh, 1)
	})
}

func onEnterVehicle(player *models.IPlayer, vehicle *models.IVehicle, seat uint8) {
	logger.LogInfof("Player %v enter vehicle: %v", player.GetName(), vehicle.GetModel())
	vehicle.SetPrimaryColor(5)
}
