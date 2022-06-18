package crud

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Product struct {
	Id    int
	Name  string
	Price int
}

var products []Product
var productId int = 0

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
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func productShow(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id, _ := strconv.Atoi(query.Get("id"))

	for _, product := range products {
		if product.Id == id {
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(product)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}

func productCreate(w http.ResponseWriter, r *http.Request) {
	var product Product

	errDecode := json.NewDecoder(r.Body).Decode(&product)
	if errDecode != nil {
		http.Error(w, errDecode.Error(), http.StatusBadRequest)
	}

	productId++
	product.Id = productId
	products = append(products, product)

	w.Header().Set("Content-Type", "application/json")
	errEncode := json.NewEncoder(w).Encode(product)
	if errEncode != nil {
		http.Error(w, errEncode.Error(), http.StatusInternalServerError)
	}
}

func productUpdate(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id, _ := strconv.Atoi(query.Get("id"))

	for index, product := range products {
		errDecode := json.NewDecoder(r.Body).Decode(&product)
		if errDecode != nil {
			http.Error(w, errDecode.Error(), http.StatusBadRequest)
		}

		if product.Id == id {
			products[index].Name = product.Name
			products[index].Price = product.Price

			_, err := w.Write([]byte("Success update data"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}

func productDelete(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	id, _ := strconv.Atoi(query.Get("id"))

	for index, product := range products {
		if product.Id == id {
			products = append(products[:index], products[index+1:]...)
			_, err := w.Write([]byte("Success delete data"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}
