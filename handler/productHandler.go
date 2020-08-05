package handlers

import (
	"encoding/json"
	"time"

	"log"
	"net/http"
	"strconv"

	helpers "../helpers"
	model "../models"
	"github.com/gorilla/mux"
)

var productStore = make(map[string]model.Product)
var id int = 0

//HTTP Post - /api/products
func PostProductHandler(w http.ResponseWriter, r *http.Request) {

	var product model.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	helpers.CheckError(err)
	
	product.CreatedOn = time.Now()
	id++
	product.ID = id
	key := strconv.Itoa(id)
	productStore[key] = product

	data, err := json.Marshal(product)
	helpers.CheckError(err)

	w.Header().Set("Content-Type", "application/json") // gelen verinin tipini belirtiyoruz
	w.WriteHeader(http.StatusCreated)
	w.Write(data)

}

//HTTP Get - /api/products
func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	var products []model.Product
	for _, product := range productStore {
		products = append(products, product)
	}

	data, err := json.Marshal(products)
	helpers.CheckError(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

//HTTP Get - /api/products/{id}
func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])
	for _, prdct := range productStore {
		if prdct.ID == key {
			product = prdct
		}
	}
	data, err := json.Marshal(product)
	helpers.CheckError(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

//HTTP Put - /api/products/{id}
func PutProductHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	key := vars["id"]

	var prodctUpd model.Product
	err = json.NewDecoder(r.Body).Decode(&prodctUpd)
	helpers.CheckError(err)

	if _, ok := productStore[key]; ok {
		prodctUpd.ID, _ = strconv.Atoi(key)
		prodctUpd.ChangedOn = time.Now()
		delete(productStore, key)
		productStore[key] = prodctUpd
	} else {
		log.Printf("Değer bulunamadı: %s", key)
	}
	w.WriteHeader(http.StatusOK)

}

//HTTP Delete -/api/products/{id}
func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	if _, ok := productStore[key]; ok {
		delete(productStore, key)

	} else {
		log.Printf("Değer bulunamadı: %s", key)
	}
	w.WriteHeader(http.StatusOK)
}
