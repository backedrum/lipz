package api

import (
	"encoding/json"
	"github.com/backedrum/lipz/models"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

// CapturePackets performs packets capture based on provided interface name and duration.
func CapturePackets(w http.ResponseWriter, req *http.Request) {
	device := mux.Vars(req)["interfaceName"]
	if device == "" {
		log.Println("Missing device name.")
		http.Error(w, "Missing device name.", http.StatusBadRequest)
		return
	}

	durationStr := mux.Vars(req)["duration"]
	duration, err := strconv.Atoi(durationStr)
	if err != nil {
		msg := "Cannot parse duration." + err.Error()
		log.Printf("Cannot parse duration %s. Error:%s.", durationStr, err)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	handle, err := pcap.OpenLive(device, 65535, false, -1*time.Second)
	defer handle.Close()
	if err != nil {
		log.Println(err)
		http.Error(w, "An error has been occurred while opening device for capture."+err.Error(), http.StatusInternalServerError)
		return
	}
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	start := time.Now()

	netPackets := []models.NetPacketInfo{}

	for packet := range packetSource.Packets() {
		if time.Now().Sub(start).Seconds() > float64(duration) {
			log.Printf("Stopping packets capture after %d seconds.", duration)
			break
		}

		netPacket := models.NetPacketInfo{}

		// set payload
		appLayer := packet.ApplicationLayer()
		if appLayer != nil {
			netPacket.Payload = string(appLayer.Payload())
		}

		// set source and destination IPs
		ip4Layer := packet.Layer(layers.LayerTypeIPv4)
		if ip4Layer != nil {
			ip, _ := ip4Layer.(*layers.IPv4)
			netPacket.Protocol = ip.Protocol.String()
			netPacket.Src = ip.SrcIP.String()
			netPacket.Dst = ip.DstIP.String()
		}

		// set source and destination ports
		tcpLayer := packet.Layer(layers.LayerTypeTCP)
		if tcpLayer != nil {
			tcp, _ := tcpLayer.(*layers.TCP)
			netPacket.SrcPort = tcp.SrcPort.String()
			netPacket.DstPort = tcp.DstPort.String()
		}

		netPackets = append(netPackets, netPacket)
	}

	js, _ := json.Marshal(models.NetPacketInfoList{netPackets})
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
