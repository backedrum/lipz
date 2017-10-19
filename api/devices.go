package api

import "net/http"
import (
	"encoding/json"
	"github.com/backedrum/lipz/models"
	"github.com/google/gopacket/pcap"
	"log"
	"sort"
)

// List available network devices
func (api *API) Devices(w http.ResponseWriter, req *http.Request) {
	var d []pcap.Interface
	d, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	devices := models.Devices{d}

	sort.Slice(devices.Devices, func(i, j int) bool {
		return devices.Devices[i].Name < devices.Devices[j].Name
	})

	js, _ := json.Marshal(devices)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
