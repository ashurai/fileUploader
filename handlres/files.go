package handlres

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"

	"github.com/ashurai/fileUploader/database"
)

func GetFiles(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(mux.Vars(r)["page"])
	if err != nil {
		w.WriteHeader(204)
		json.NewEncoder(w).Encode("wrong value passed for page paramater")
		return
	}
	limit := 5 // Best way to define globaly or use a constat , 
	// but this is only one place in current implementation
	var offset int

	if page > 1 {
		offset = (limit * (page - 1)) // 5 * (3 - 1) = 10
	} else {
		offset = 0
	}

	fiels := database.GetFiles(offset, limit)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
    //it returns the new file details
    json.NewEncoder(w).Encode(fiels)
}