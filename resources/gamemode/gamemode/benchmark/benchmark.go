package benchmark

import (
	"fmt"
	"github.com/StanZzzz222/RAltGo/common/alt/blip"
	"github.com/StanZzzz222/RAltGo/common/alt/ped"
	"github.com/StanZzzz222/RAltGo/common/alt/vehicle"
	"github.com/StanZzzz222/RAltGo/common/utils"
	"github.com/StanZzzz222/RAltGo/hash_enums"
	"github.com/StanZzzz222/RAltGo/hash_enums/ped_hash"
	"github.com/StanZzzz222/RAltGo/hash_enums/vehicle_hash"
	"github.com/StanZzzz222/RAltGo/logger"
	"time"
)

/*
   Create by zyx
   Date Time: 2024/9/22
   File: benchmark.go
*/

/*
Benchmark
Basic API setting and acquisition performance testing
We found through rigorous testing that its performance is really good.
You can try the benchmark below. The following is the time taken by Intel Core i913900H
*/
func Benchmark() {
	ped1 := ped.CreatePedByHash(ped_hash.Abner, utils.NewVector3(-1019.3187, -2928.9758, 14), utils.NewVector3(0, 0, 0))
	start := time.Now()
	for i := 0; i < 50000; i++ {
		_ = vehicle.CreateVehicleByHash(vehicle_hash.T20, fmt.Sprintf("test%v", i), utils.NewVector3(-1069.3187, -2928.9758, 14.1318), utils.NewVector3(0, 0, 0), 1, 1)
	}
	logger.LogInfof("Create 50,000 vehicles Since: %v ms", time.Since(start).Milliseconds()) // Since: 617 ms
	start = time.Now()
	for i := 0; i < 10000; i++ {
		_ = blip.CreateBlipPoint(12, 1, fmt.Sprintf("test%v", i), utils.NewVector3(-1069.3187, -2928.9758, 14.1318))
		_ = blip.CreateBlipRadius(1, fmt.Sprintf("test%v", i), utils.NewVector3(-1069.3187, -2928.9758, 14.1318), 15, false)
		_ = blip.CreateBlipArea(1, fmt.Sprintf("test%v", i), utils.NewVector3(-1069.3187, -2928.9758, 14.1318), 15, 15)
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
