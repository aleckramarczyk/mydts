package utils

import (
	"aleckramarczyk/mydts/agent/entities"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func SendInfoRequest(Mdt *entities.MDT) error {
	url := getEndpoint()
	body, err := json.Marshal(Mdt)
	if err != nil {
		log.Printf("Error encoding MDT object: %s\n", err)
		return err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Error creating POST request: %s\n", err)
		return err
	}
	defer resp.Body.Close()
	return nil
}

func getEndpoint() string {
	return fmt.Sprintf("http://%s:%s", AgentConfig.API_Host, AgentConfig.API_Endpoint)
}
