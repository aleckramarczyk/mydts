package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"aleckramarczyk/mydts/server/db"
	"aleckramarczyk/mydts/server/entities"
	"aleckramarczyk/mydts/server/utils"
)

func CreateMDT(mdt *entities.MDT) {
	if utils.AppConfig.DEBUG {
		log.Println("createMDT triggered")
	}
	db.Instance.Create(&mdt)
	log.Printf("CREATE | ID: %s | REMOTE_IP: %s | DOCK_MAC: %s", mdt.Mdt_uuid, mdt.Remote_ip, mdt.Dock_mac)
}

func UpdateMDT(newMdt *entities.MDT) {
	if utils.AppConfig.DEBUG {
		log.Println("updateMDT triggered")
	}
	var oldMdt entities.MDT
	db.Instance.First(&oldMdt, newMdt.Dock_mac)
	oldMdt.Dock_mac = newMdt.Dock_mac
	oldMdt.Mdt_uuid = newMdt.Mdt_uuid
	oldMdt.Remote_ip = newMdt.Remote_ip
	oldMdt.Updated = newMdt.Updated
	db.Instance.Save(&oldMdt)
	log.Printf("UPDATE | ID: %s | REMOTE_IP: %s | DOCK_MAC: %s", oldMdt.Mdt_uuid, oldMdt.Remote_ip, oldMdt.Dock_mac)
}

func postMDT(w http.ResponseWriter, r *http.Request) {
	var mdt entities.MDT
	json.NewDecoder(r.Body).Decode(&mdt)
	mdt.Remote_ip = r.RemoteAddr
	mdt.Updated = time.Now()

	if db.MDTExists(mdt.Dock_mac) {
		UpdateMDT(&mdt)
		if utils.AppConfig.DEBUG {
			log.Printf("DEBUG: Post request. mac: %s, mdtExists: true\n", mdt.Dock_mac)
		}
	} else {
		CreateMDT(&mdt)
		if utils.AppConfig.DEBUG {
			log.Printf("DEBUG: Post request. Mac: %s, mdtExists: false\n", mdt.Dock_mac)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mdt)
}

func getMDTs(w http.ResponseWriter, r *http.Request) {
	var mdts []entities.MDT
	db.Instance.Find(&mdts)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mdts)
}
