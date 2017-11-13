package models

type NetPacketInfo struct {
	Protocol string
	Src      string
	SrcPort  string
	Dst      string
	DstPort  string
	Payload  string
}

type NetPacketInfoList struct {
	NetPacketInfoList []NetPacketInfo
}
