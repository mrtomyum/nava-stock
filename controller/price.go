package controller

import (
	"net/http"
	"github.com/gorilla/mux"
	m "github.com/mrtomyum/nava-stock/model"
	"strconv"
	"log"
	"github.com/mrtomyum/nava-stock/api"
	"encoding/json"
	"fmt"
)

func (e *Env) ItemPrice(w http.ResponseWriter, r *http.Request) {
	log.Println("call ItemPrice")
	v := mux.Vars(r)
	id := v["id"]
	ip := new(m.ItemPrice)
	ips := []*m.ItemPrice{}
	ip.ItemID, _ = strconv.ParseUint(id, 10, 64)
	log.Println("item.ID = ", ip.ItemID)

	ips, err := ip.AllPrice(e.DB)

	rs := new(api.Response)
	if err != nil {
		rs.Status = "204"
		rs.Message = "NO CONTENT"
		rs.Result = nil
	} else {
		rs.Status = "200"
		rs.Message = "SUCCESS return All Item Price "
		rs.Result = ips
	}
	o, _ := json.Marshal(rs)
	fmt.Fprintf(w, string(o))
}

