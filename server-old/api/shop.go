package api

import (
	"aleckramarczyk/mydts/server/db"
	"aleckramarczyk/mydts/server/entities"
	"aleckramarczyk/mydts/server/utils"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func postShop(w http.ResponseWriter, r *http.Request) {
	var shop entities.Shop
	err := json.NewDecoder(r.Body).Decode(&shop)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if db.ShopExists(shop.Dock_mac) {
		updateShop(&shop)
	} else {
		createShop(&shop)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(shop)
}

func createShop(shop *entities.Shop) {
	if utils.AppConfig.DEBUG {
		log.Println("createShop triggered")
	}
	db.Instance.Create(&shop)
	log.Printf("CREATE SHOP | SHOP_NUMBER: %s | DOCK_MAC: %s", shop.Shop_number, shop.Dock_mac)
}

func updateShop(newShop *entities.Shop) {
	if utils.AppConfig.DEBUG {
		log.Println("updateShop triggered")
	}
	var oldShop entities.Shop
	db.Instance.First(&oldShop, newShop)
	oldShop.Dock_mac = newShop.Dock_mac
	oldShop.Shop_number = newShop.Shop_number
	db.Instance.Save(&oldShop)
	log.Printf("UPDATE SHOP | DOCK_MAC: %s | SHOP_NUMBER: %s", oldShop.Dock_mac, oldShop.Shop_number)
}

func getShops(w http.ResponseWriter, r *http.Request) {
	var shops []entities.Shop
	db.Instance.Find(&shops)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(shops)
}

func deleteShop(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	shopNumber := mux.Vars(r)["shopNumber"]
	if db.ShopExistsByShopNumber(shopNumber) {
		var shop entities.Shop
		db.Instance.Delete(&shop, shopNumber)
		json.NewEncoder(w).Encode("Shop deleted successfully")
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Shop not found")
	}
}

func getShopByShopNumber(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	shopNumber := mux.Vars(r)["shopNumber"]
	if db.ShopExistsByShopNumber(shopNumber) {
		var shop entities.Shop
		db.Instance.First(&shop, shopNumber)
		json.NewEncoder(w).Encode(shop)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Shop not found")
	}
}
