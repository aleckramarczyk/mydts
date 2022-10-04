package main

import (
	"aleckramarczyk/mydts/agent/entities"
	"aleckramarczyk/mydts/agent/utils"
)

var mdt *entities.MDT

func main() {
	//get config information
	utils.LoadConfig()
	//gather information
	mdt.Dock_MAC = utils.GetMAC()
	mdt.MDT_UUID = utils.GET_UUID()

	//Send request to API endpoint

	//sleep and repeat
}
