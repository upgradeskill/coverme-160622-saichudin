package crud

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/saichudin/golang-starter/model/web"
	"github.com/saichudin/golang-starter/service"
)

func CrudBasic() {
	http.HandleFunc("/products/", handlerProducts)
	http.HandleFunc("/product/", handlerProduct)

	server := http.Server{
		Addr: "localhost:8080",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func handlerProducts(w http.ResponseWriter, r *http.Request) {
	productIndex(w, r)
}

func handlerProduct(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		productShow(w, r)
	case "POST":
		productCreate(w, r)
	case "PUT":
		productUpdate(w, r)
	case "DELETE":
		productDelete(w, r)
	default:
		http.Error(w, "", http.StatusBadRequest)
	}
}

func productIndex(w http.ResponseWriter, r *http.Request) {
	productService := service.NewProductService()
	productResp := productService.FindAll()
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(productResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func productShow(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id, _ := strconv.Atoi(query.Get("id"))

	productService := service.NewProductService()
	productResp := productService.FindById(id)

	w.Header().Set("Content-Type", "application/json")
	errEncode := json.NewEncoder(w).Encode(productResp)
	if errEncode != nil {
		http.Error(w, errEncode.Error(), http.StatusInternalServerError)
	}
}

func productCreate(w http.ResponseWriter, r *http.Request) {
	productService := service.NewProductService()
	decoder := json.NewDecoder(r.Body)
	productRequest := web.ProductRequest{}
	errDecode := decoder.Decode(&productRequest)
	if errDecode != nil {
		http.Error(w, errDecode.Error(), http.StatusBadRequest)
	}
	productResp := productService.Create(productRequest)

	w.Header().Set("Content-Type", "application/json")
	errEncode := json.NewEncoder(w).Encode(productResp)
	if errEncode != nil {
		http.Error(w, errEncode.Error(), http.StatusInternalServerError)
	}
}

func productUpdate(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id, _ := strconv.Atoi(query.Get("id"))

	productService := service.NewProductService()
	decoder := json.NewDecoder(r.Body)
	productRequest := web.ProductRequest{}
	errDecode := decoder.Decode(&productRequest)
	if errDecode != nil {
		http.Error(w, errDecode.Error(), http.StatusBadRequest)
	}

	productResp := productService.Update(id, productRequest)

	w.Header().Set("Content-Type", "application/json")
	errEncode := json.NewEncoder(w).Encode(productResp)
	if errEncode != nil {
		http.Error(w, errEncode.Error(), http.StatusInternalServerError)
	}
}

func productDelete(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id, _ := strconv.Atoi(query.Get("id"))

	productService := service.NewProductService()
	productService.Delete(id)

	_, err := w.Write([]byte("Success delete data"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
