package api

import (
	"encoding/json"
	"github.com/backedrum/lipz/models"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
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

	var cs models.CaptureSettings
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&cs); err != nil {
		msg := "Cannot parse capture settings" + err.Error()
		log.Printf("%s, request body:%s", msg, req.Body)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	var pw *pcapgo.Writer
	if cs.Filename != "undefined" {
		log.Printf("Received filename: %s", cs.Filename)

		f, _ := os.Create(cs.Filename)
		defer f.Close()

		pw = pcapgo.NewWriter(f)
		pw.WriteFileHeader(65535, 1)
	}

	handle, err := pcap.OpenLive(device, 65535, false, -1*time.Second)
	defer handle.Close()
	if err != nil {
		log.Println(err)
		http.Error(w, "An error has been occurred while opening device for capture."+err.Error(), http.StatusInternalServerError)
		return
	}
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	netPackets := []models.NetPacketInfo{}

	// perform capture for the specified duration
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(time.Duration(cs.Duration) * time.Second)
		timeout <- true
	}()

capture:
	for packet := range packetSource.Packets() {
		select {
		case <-timeout:
			log.Printf("Stopping packets capture after %d seconds.", cs.Duration)
			break capture
		default:
			break
		}

		// write captured packet to file
		if pw != nil {
			pw.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
		}

		netPacket := models.NetPacketInfo{}


		appLayer := packet.ApplicationLayer()
		if appLayer != nil {
			// set dump
			netPacket.Dump = packet.Dump()

			// set payload
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
