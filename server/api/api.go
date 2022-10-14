package api

import "github.com/gorilla/mux"

func RegisterApiRoutes(router *mux.Router) {
	router.HandleFunc("/api/mdt", postMDT).Methods("POST")
	router.HandleFunc("/api/shop", postShop).Methods("POST")
	router.HandleFunc("/api/mdt", getMDTs).Methods("GET")
	router.HandleFunc("/api/shop", getShops).Methods("GET")
	router.HandleFunc("/api/shop/{shopNumber}", deleteShop).Methods("DELETE")
	router.HandleFunc("/api/shop/{shopNumber}", getShopByShopNumber).Methods("GET")
}
