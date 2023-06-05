package utils

import (
	"fmt"
	"github.com/jackpal/gateway"
	"log"
	"net"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

func GetSerialNumber() (string, error) {
	out, err := exec.Command("cmd", "/C", "wmic bios get serialnumber | findstr /v SerialNumber").Output()
	if err != nil {
		log.Printf("Error getting UUID: %s\n", err)
		return "", err
	}
	convertOut := strings.TrimSpace(string(out))
	return convertOut, nil
}

func SendUpdateOnInterval() {
	for {
		time.Sleep(AgentConfig.UpdateInterval)
		err := SendUpdateRequest(UnitInformation)
		if err != nil {
			log.Printf("Error sending update request: %s. Skipping update", err)
			continue
		}
	}
}

/*
func GetInterfacesHardwareAddrs() (map[string]string, error) {
	interfaces := make(map[string]string)
	intrs, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, intr := range intrs {
		interfaces[intr.Name] = intr.HardwareAddr.String()
	}
	return interfaces, nil
}
*/

/*
func ConnectedToDock() (bool, error) {
	netInters, err := net.Interfaces()
	if err != nil {
		log.Printf("Error getting network interfaces: %s\n", err.Error())
		return false, err
	}
	for _, netInter := range netInters {
		if netInter.Name == "Ethernet 3" || netInter.Name == "Ethernet 4" {
			return true, nil
		}
	}
	return false, nil
}
*/

func GetDefaultGatewayMac() (gatewayMac string, err error) {
	gatewayIp, err := getDefaultGateway()
	if err != nil {
		return "", err
	}
	out, err := exec.Command("cmd", "/C", "arp -a "+gatewayIp).Output()
	macLine := strings.Split(string(out), "\n")[3]
	re := regexp.MustCompile(`\s+`)
	macLine = re.ReplaceAllString(macLine, " ")
	gatewayMac = strings.TrimSpace(strings.Split(macLine, " ")[2])
	fmt.Println(gatewayMac)
	return gatewayMac, err
}

func getDefaultGateway() (string, error) {
	gatewayIp, err := gateway.DiscoverGateway()
	if err != nil {
		log.Printf("Error getting default gateway: %s\n", err.Error())
	}
	return gatewayIp.String(), err
}

func GetLocalIP() string {
	// This function does not actually send any traffic, it is only detecting the IP of the default interface that the machine would use to send this traffic
	conn, err := net.Dial("udp", "192.168.1.0:80")
	if err != nil {
		log.Println(err.Error())
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}

/*
// ConvertGPRMC receives a GPRMC gps format as a slice of strings and returns the latitude and longitude it contains.
// Intergraph generates gps data in this format, and it needs to be converted to the more familiar lon lat
func ConvertGPRMC(gprmc []string) (utc int, latitude float64, longitude float64, err error) {
	if gprmc[0] != "$GPRMC" {
		return 0, 0, 0, errors.New("invalid GPRMC format")
	}
	utc, _ = strconv.Atoi(gprmc[1])
	nmeaLat := gprmc[3]
	latQuadrant := rune((gprmc[4])[0])
	nmeaLon := gprmc[5]
	lonQuadrant := rune((gprmc[6])[0])

	latitude = GpsToDecimalDegrees(nmeaLat, latQuadrant)
	longitude = GpsToDecimalDegrees(nmeaLon, lonQuadrant)

	return utc, longitude, latitude, nil
}

// GpsToDecimalDegrees converts NMEA absolute position to decimal degrees. NMEA format is ddmm.mmm... n/s, (d)ddmm.mmm... e/w
// the formula to convert to decimal degrees is (d)dd + (mm.mmm.../60) * (-1 if quadrant = s || w)
func GpsToDecimalDegrees(nmeaPos string, quadrant rune) (decimal float64) {
	dotIndex := strings.Index(nmeaPos, ".")
	if dotIndex > -1 {
		mmIndex := dotIndex - 2
		dd, _ := strconv.ParseFloat(nmeaPos[0:mmIndex], 64)
		mm, _ := strconv.ParseFloat(nmeaPos[mmIndex:], 64)
		decimal = dd + (mm / 60)
		if quadrant == 'S' || quadrant == 's' || quadrant == 'W' || quadrant == 'w' {
			decimal *= -1
		}
	}
	return decimal
}
*/
