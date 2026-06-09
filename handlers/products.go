package handlers

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	InStock  bool    `json:"in_stock"`
}

var products = []Product{
	{ID: "1", Name: "Widget A", Price: 9.99, InStock: true},
	{ID: "2", Name: "Widget B", Price: 24.99, InStock: false},
	{ID: "3", Name: "Gadget X", Price: 49.99, InStock: true},
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
