package controller

import (
	m "github.com/mrtomyum/nava-stock/model"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"fmt"
	"encoding/json"
	"strconv"
	"github.com/mrtomyum/nava-api3/api"
)

func CreateLocationTree(locations []*m.Location) *m.Location {
	tree := new(m.Location)
	for _, l := range locations {
		tree.Add(l)
	}
	return tree
}

func (e *Env) ShowLocationTree(w http.ResponseWriter, r *http.Request) {
	log.Println("call ShowLocationTree()")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	loc := new(m.Location)
	v := mux.Vars(r)
	id := v["id"]
	loc.ID, _ = strconv.ParseUint(id, 10, 64)
	// Todo: use loc.ID to parameter to retrive just tree of this ID
	locations, err := loc.All(e.DB)
	rs := new(api.Response)
	if err != nil {
		log.Fatal("ShowLocations()", err)
		w.WriteHeader(http.StatusNotFound)
		rs.Status = api.ERROR
		rs.Message = "Location not found or Error."
	} else {
		w.WriteHeader(http.StatusOK)
		tree := CreateLocationTree(locations)
		rs.Status = api.SUCCESS
		rs.Data = tree
	}
	output, _ := json.Marshal(rs)
	fmt.Fprintf(w, string(output))
}
