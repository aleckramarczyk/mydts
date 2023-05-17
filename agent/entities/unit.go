package entities

type Unit struct {
	UnitName     string `json:"unitName"`
	DockMac      string `json:"dockMac"`
	MdtUuid      string `json:"mdtUuid"`
	UnitId       string `json:"unitId"`
	InternalIP   string `json:"internalIP"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	GpsTimestamp string `json:"gpsTimestamp"`
}
