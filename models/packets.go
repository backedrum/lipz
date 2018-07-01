package models

type NetPacketInfo struct {
	Protocol string
	Src      string
	SrcPort  string
	Dst      string
	DstPort  string
	Dump     string
	Payload  string
}

type NetPacketInfoList struct {
	NetPacketInfoList []NetPacketInfo
}

type CaptureSettings struct {
	Duration int    `json:"duration"`
	Filename string `json:"filename"`
}
