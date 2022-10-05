package utils

import (
	"log"
	"net"
	"os/exec"
	"strings"
)

func GET_UUID() (string, error) {
	out, err := exec.Command("cmd", "/C", "wmic csproduct get uuid | findstr /v UUID").Output()
	if err != nil {
		log.Printf("Error getting UUID: %s\n", err)
		return "", err
	}
	convertOut := strings.TrimSpace(string(out))
	return convertOut, nil
}

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

func ConnectedToDock(interfaces map[string]string) bool {
	if _, ok := interfaces["Ethernet 3"]; ok {
		return true
	}
	return false
}
