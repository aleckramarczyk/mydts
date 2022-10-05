package main

import (
	"aleckramarczyk/mydts/agent/entities"
	"aleckramarczyk/mydts/agent/utils"
	"log"
	"time"
)

func main() {
	//get config information
	utils.LoadConfig()

	var err error
	Mdt := &entities.MDT{}

	Mdt.MDT_UUID, err = utils.GET_UUID()
	if err != nil {
		log.Fatal(err)
	}

	for {
		interfaces, err := utils.GetInterfacesHardwareAddrs()
		if err != nil {
			time.Sleep(time.Minute * 10)
			continue
		}
		if utils.ConnectedToDock(interfaces) {
			//send request
			Mdt.Dock_MAC = interfaces["Ethernet 3"]
			err := utils.SendInfoRequest(Mdt)
			if err != nil {
				time.Sleep(time.Minute * 10)
				continue
			}
		} else {
			log.Println("Not connected to dock. Skipping update")
			time.Sleep(time.Minute * 10)
			continue
		}
		log.Println("Info update successful")
		time.Sleep(time.Minute * 10)
	}
}
