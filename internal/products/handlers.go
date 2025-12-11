// This contains all the routes related to products
package products

import (
	"log"
	"net/http"

	"github.com/PulinduVR/ecom-go/internal/json"
)

type handler struct {
	service Service
}

// this is like a constructor.
func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	// 1. Call the Service -> ListProducts
	// 2. Return JSON in an HTTP response

	err := h.service.ListProducts(r.Context())

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	products := struct {
		Products []string `json:"products"`
	}{}
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(products)

	// All these things have been contained on the internal/json
	json.Write(w, http.StatusOK, products)
}
