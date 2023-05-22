package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func SendUpdateRequest(u *Unit) error {
	log.Printf("Sending update request: %s\n", time.Now().String())
	url := getEndpoint()
	body, err := json.Marshal(u)
	if err != nil {
		log.Printf("Error encoding u object: %s\n", err)
		return err
	}
	log.Println(string(body))
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Error creating POST request: %s\n", err)
		return err
	}
	defer resp.Body.Close()
	return nil
}

func getEndpoint() string {
	return fmt.Sprintf("http://%s/%s", AgentConfig.ApiHost, AgentConfig.ApiEndpoint)
}
