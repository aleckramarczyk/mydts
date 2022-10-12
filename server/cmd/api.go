package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func registerApiRoutes(router *mux.Router) {
	router.HandleFunc("/api/mdt", postMDT).Methods("POST")
}

func createMDT(mdt *MDT) {
	if AppConfig.DEBUG {
		log.Println("createMDT triggered")
	}
	Instance.Create(&mdt)
	log.Printf("CREATE | ID: %s | REMOTE_IP: %s | DOCK_MAC: %s", mdt.Mdt_uuid, mdt.Remote_ip, mdt.Dock_mac)
}

func updateMDT(newMdt *MDT) {
	if AppConfig.DEBUG {
		log.Println("updateMDT triggered")
	}
	var oldMdt MDT
	Instance.First(&oldMdt, newMdt.Dock_mac)
	oldMdt.Dock_mac = newMdt.Dock_mac
	oldMdt.Mdt_uuid = newMdt.Mdt_uuid
	oldMdt.Remote_ip = newMdt.Remote_ip
	oldMdt.Updated = newMdt.Updated
	Instance.Save(&oldMdt)
	log.Printf("UPDATE | ID: %s | REMOTE_IP: %s | DOCK_MAC: %s", oldMdt.Mdt_uuid, oldMdt.Remote_ip, oldMdt.Dock_mac)
}

func postMDT(w http.ResponseWriter, r *http.Request) {
	var mdt MDT
	json.NewDecoder(r.Body).Decode(&mdt)
	mdt.Remote_ip = r.RemoteAddr
	mdt.Updated = time.Now()

	if MDTExists(mdt.Dock_mac) {
		updateMDT(&mdt)
		if AppConfig.DEBUG {
			log.Printf("DEBUG: Post request. mac: %s, mdtExists: true\n", mdt.Dock_mac)
		}
	} else {
		createMDT(&mdt)
		if AppConfig.DEBUG {
			log.Printf("DEBUG: Post request. Mac: %s, mdtExists: false\n", mdt.Dock_mac)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mdt)
}
