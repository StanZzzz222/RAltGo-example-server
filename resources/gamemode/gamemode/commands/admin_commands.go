package commands

import (
	"fmt"
	"github.com/StanZzzz222/RAltGo/common/alt/command"
	"github.com/StanZzzz222/RAltGo/common/alt/vehicle"
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/hash_enums/vehicle_hash"
	"github.com/StanZzzz222/RAltGo/logger"
)

/*
   Create by zyx
   Date Time: 2024/9/22
   File: admin_commands.go
*/

func InitAdminCommands() {
	group := command.NewCommandGroup("AdminCommands")
	// Do logic processing in middleware?
	group.UseMiddleware(func(player *models.IPlayer, name string, args []any) bool {
		if player.HasData("admin") {
			logger.LogInfof("Player: %v use command: %v", player.GetName(), name)
			return true
		}
		player.SendBroadcast("You do not have permission to use this command")
		return false
	})
	// Add OnCommand
	{
		group.OnCommandDesc("createveh", CreateVehicle, false, "/createveh [VehicleName(example: t20)]")
		group.OnCommandDesc("setvehcolor", SetVehicleColor, false, "/setvehcolor [PrimaryColor(1-159)] [SecondColor(1-159)]")
	}
}

func CreateVehicle(player *models.IPlayer, name string) {
	if player.HasData("veh") {
		veh := player.GetData("veh").(*models.IVehicle)
		veh.Destroy()
	}
	veh := vehicle.CreateVehicle(name, "test", player.GetPosition(), player.GetRotation(), 1, 1)
	player.SetData("veh", veh)
	player.SetIntoVehicle(veh, uint8(vehicle_hash.Driver))
	player.SendBroadcast(fmt.Sprintf("Create vehicle: %v | id: %v", veh.GetModel().String(), veh.GetId()))
}

func SetVehicleColor(player *models.IPlayer, primaryColor, secondColor int64) {
	if (primaryColor <= 0 || secondColor <= 0) || (primaryColor >= 159 && secondColor >= 159) {
		player.SendBroadcast(fmt.Sprintf("PrimaryColor or SecondColor range in 1-159"))
		return
	}
	if !player.IsInVehicle() {
		player.SendBroadcast(fmt.Sprintf("You are not currently in any vehicle"))
		return
	}
	veh := player.Vehicle()
	veh.SetPrimaryColor(uint8(primaryColor))
	veh.SetSecondColor(uint8(secondColor))
}