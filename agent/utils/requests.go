package utils

import (
	"aleckramarczyk/mydts/agent/entities"
	"fmt"
)

func SendInfoRequest(mdt *entities.MDT) error {
	return nil
}

func getEndpoint() string {
	return fmt.Sprintf("http://%s:%s", AgentConfig.API_Host, AgentConfig.API_Endpoint)
}
