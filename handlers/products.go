package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/Ubivius/microservice-template/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// If we don't handle specific methods, the same function is called for get, post, delete, put
func (p *Products) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	// Handling get
	if request.Method == http.MethodGet {
		p.getProducts(w, request)
		return
	}

	if request.Method == http.MethodPost {
		p.addProduct(w, request)
		return
	}

	if request.Method == http.MethodPut {
		p.l.Println("PUT", request.URL.Path)
		// Expect the id in the URI
		regex := regexp.MustCompile(`/([0-9]+)`)
		group := regex.FindAllStringSubmatch(request.URL.Path, -1)

		if len(group) != 1 {
			p.l.Println("Invalide URI, more than one id")
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(group[0]) != 2 {
			p.l.Println("Invalide URI, more than one capture group")
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}

		// Extract the id from the regex result
		idString := group[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.l.Println("Valid URI, unable to convert id to int")
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}

		p.updateProducts(id, w, request)
		return
	}

	// If method is not implemented, return an error
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) updateProducts(id int, w http.ResponseWriter, request *http.Request) {
	p.l.Println("Handle PUT product")
	product := &data.Product{}

	// Parsing to json
	err := product.FromJson(request.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
	}

	// Update product
	err = data.UpdateProduct(id, product)
	if err == data.ErrorProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
}

func (p *Products) addProduct(w http.ResponseWriter, request *http.Request) {
	p.l.Println("Handle POST product")
	product := &data.Product{}
	err := product.FromJson(request.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
	}
	data.AddProduct(product)
}

func (p *Products) getProducts(w http.ResponseWriter, request *http.Request) {
	p.l.Println("Handle GET products")
	productList := data.GetProducts()
	err := productList.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}
