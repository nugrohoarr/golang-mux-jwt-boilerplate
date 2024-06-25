package productcontroller

import (
	"net/http"

	"github.com/nugrohoarr/golang-mux-jwt-boilerplate/helper"
)

func Index(w http.ResponseWriter, r *http.Request) {

	data := []map[string]interface{}{
		{
			"id":           1,
			"product_name": "Product 1",
			"price":        100000,
		},
		{
			"id":           2,
			"product_name": "Product 2",
			"price":        10000,
		},
		{
			"id":           1,
			"product_name": "Product 1",
			"price":        500000,
		},
	}

	helper.ResponseJSON(w, http.StatusOK, data)
}
