package utils

import "log"

type Unit struct {
	UnitName   string `json:"unitName"`
	UnitId     string `json:"unitId"`
	VehicleID  string `json:"vehicleID"`
	SignedOn   bool   `json:"signedOn"`
	InternalIP string `json:"internalIP"`
}

var UnitInformation = new(Unit)
var MDTSerialNumber string

/*
func (u *Unit) UpdateGPSInformation(e string) {
	lat, lon, err := utils.ReadGPSInformation(e)
	if err != nil {
		log.Println("ERROR reading GPS info:", err)
		return
	}

	// This shouldn't happen as the sign on event should be triggered before a GPS update,
	// but to prevent errors in case it does we can read UnitID, UnitName, and VehicleID
	// from the ApplicationState.xml file
	if u.UnitName == "" {
		err = utils.ReadApplicationState(u)
		if err != nil {
			log.Println("ERROR reading application state:", err)
			return
		}
	}

	u.Latitude = lat
	u.Longitude = lon
	u.GpsTimestamp = time.Now() // Could be parsed from the GPRMC gps data for a little bit more accuracy, but this is much easier
}
*/

func (u *Unit) UpdateUnitPropertiesOnSignOn(e string) {
	err := ReadAuthorizeInformation(u, e)
	if err != nil {
		log.Println("ERROR reading authorization information:", err)
	}

	//Get local IP before update in case it has changed
	u.InternalIP = GetLocalIP()

	err = SendUpdateRequest(u)
	if err != nil {
		log.Println("ERROR sending authorize request:", err)
	}
}

func (u *Unit) UpdateUnitPropertiesOnSignOff(e string) {
	// Again, this shouldn't happen. But here we are
	if u.UnitName == "" {
		err := ReadApplicationState(u)
		if err != nil {
			log.Println("ERROR reading application state:", err)
			return
		}
	}

	// Sign off archive's don't contain much information, so we just go with previously collected information
	u.SignedOn = false

	err := SendUpdateRequest(u)
	if err != nil {
		log.Println("ERROR sending sign off request:", err)
		return
	}
}
