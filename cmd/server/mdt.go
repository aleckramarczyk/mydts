package main

import "time"

type MDT struct {
	ID      string    `json:"id"`
	DockMac string    `json:"dock_mac"`
	IP      string    `json:"ip"`
	Updated time.Time `json:"updated"`
}
