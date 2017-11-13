package api

import (
	"encoding/json"
	"fmt"
	"github.com/backedrum/lipz/models"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func Packets(w http.ResponseWriter, req *http.Request) {
	device := mux.Vars(req)["interfaceName"]

	handle, err := pcap.OpenLive(device, 65535, false, -1*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	start := time.Now()

	netPackets := []models.NetPacketInfo{}

	for packet := range packetSource.Packets() {
		fmt.Printf("Difference %f\n", time.Now().Sub(start).Seconds())
		if time.Now().Sub(start).Seconds() > 15.0 {
			log.Print("Stopping packets capture after 15 seconds.")
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
