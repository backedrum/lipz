package api

import "net/http"
import (
	"encoding/json"
	"github.com/backedrum/lipz/models"
	"github.com/akrennmair/gopcap"
	"log"
)

// List available network devices
func (api *API) Devices(w http.ResponseWriter, req *http.Request) {
	var d []pcap.Interface
	d, err := pcap.Findalldevs()
	if err != nil {
		log.Fatal(err)
	}

	devices := models.Devices{d}

	js, _ := json.Marshal(devices)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}