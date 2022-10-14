package main

import (
	"aleckramarczyk/mydts/agent/entities"
	"aleckramarczyk/mydts/agent/utils"
	"log"
	"time"

	"github.com/kardianos/service"
)

const serviceName = "MDT Agent"
const serviceDescription = "Telemetry service for MDTs"

type program struct{}

func (p program) Start(s service.Service) error {
	log.Println(s.String() + " started")
	go p.run()
	return nil
}

func (p program) Stop(s service.Service) error {
	log.Println(s.String() + " stopped")
	return nil
}

func (p program) run() {
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

func main() {
	//get config information
	utils.LoadConfig()

	serviceConfig := &service.Config{
		Name:        serviceName,
		DisplayName: serviceName,
		Description: serviceDescription,
	}

	prg := &program{}
	s, err := service.New(prg, serviceConfig)
	if err != nil {
		log.Println("Cannot create the service " + err.Error())
	}
	err = s.Run()
	if err != nil {
		log.Println("Cannot start the service " + err.Error())
	}

}
