package api

import "net/http"
import (
	"encoding/json"
	"fmt"
	"github.com/backedrum/lipz/models"
	"github.com/google/gopacket/pcap"
	"log"
	"sort"
)

// List available network devices
func Devices(w http.ResponseWriter, req *http.Request) {
	var d []pcap.Interface
	d, err := pcap.FindAllDevs()
	if err != nil {
		msg := fmt.Sprintf("An error has been occurred while retrieving devices list: %s", err.Error())
		log.Println(msg)
		http.Error(w, "Cannot retrieve devices list.", http.StatusInternalServerError)
		return
	}

	devices := models.Devices{d}

	sort.Slice(devices.Devices, func(i, j int) bool {
		return devices.Devices[i].Name < devices.Devices[j].Name
	})

	js, _ := json.Marshal(devices)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
