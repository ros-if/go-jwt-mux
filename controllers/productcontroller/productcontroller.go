package productcontroller

import (
	"net/http"

	"github.com/ros-if/jwt/helper"
)

func Index(w http.ResponseWriter, r *http.Request) {

	data := []map[string]interface{}{
		{
			"id": 1,
			"nama_product": "Buku 1",
			"ukuran": "A3",
		},
		{
			"id": 2,
			"nama_product": "Buku 2",
			"ukuran": "A4",
		},
		{
			"id": 3,
			"nama_product": "Buku 3",
			"ukuran": "A5",
		},
	}

	helper.ResponseJSON(w, http.StatusOK, data)
}