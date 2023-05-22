package main

import (
	"aleckramarczyk/mydts/agent/utils"
	"github.com/kardianos/service"
	"log"
)

const serviceName = "Unit Agent"
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
	go utils.WatchForEvents()

	if utils.AgentConfig.UpdateOnInterval {
		utils.SendUpdateOnInterval()
	}
}

func init() {
	var err error

	utils.LoadConfig()

	// Get initial information about the unit
	utils.MDTSerialNumber, err = utils.GetSerialNumber()
	if err != nil {
		log.Fatal("Fatal error getting UUID:", err)
	}
	utils.UnitInformation.InternalIP = utils.GetLocalIP()

}

func main() {
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
