package models

import "github.com/google/gopacket/pcap"

type Devices struct {
	Devices []pcap.Interface
}
