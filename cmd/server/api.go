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

func createMDT(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var mdt MDT
	json.NewDecoder(r.Body).Decode(&mdt)
	mdt.IP = r.RemoteAddr
	mdt.Updated = time.Now()
	Instance.Create(&mdt)
	json.NewEncoder(w).Encode(mdt)
	log.Printf("CREATE | ID: %s | REMOTE_IP: %s | DOCK_MAC: %s", mdt.ID, mdt.IP, mdt.DockMac)
}

func updateMDT(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var mdt MDT
	Instance.First(&mdt, id)
	json.NewDecoder(r.Body).Decode(&mdt)
	mdt.IP = r.RemoteAddr
	mdt.Updated = time.Now()
	Instance.Save(&mdt)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mdt)
	log.Printf("UPDATE | ID: %s | REMOTE_IP: %s | DOCK_MAC: %s", mdt.ID, mdt.IP, mdt.DockMac)
}

func postMDT(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if MDTExists(id) {
		updateMDT(w, r)
		return
	}
	createMDT(w, r)
}
