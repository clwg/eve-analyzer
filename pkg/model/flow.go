package model

import (
	"github.com/clwg/eve-analyzer/types"
)

// Note: UDP has no other fields outside of the flow
type Flow struct {
	PktsToserver  int              `json:"pkts_toserver"`
	PktsToclient  int              `json:"pkts_toclient"`
	BytesToserver int              `json:"bytes_toserver"`
	BytesToclient int              `json:"bytes_toclient"`
	Start         types.CustomTime `json:"start"`
	End           types.CustomTime `json:"end"`
	Age           int              `json:"age"`
	State         string           `json:"state"`
	Reason        string           `json:"reason"`
	Alerted       bool             `json:"alerted"`
}

type TCP struct {
	TCPFlags     string `json:"tcp_flags"`
	TCPFlagsTs   string `json:"tcp_flags_ts"`
	TCPFlagsTc   string `json:"tcp_flags_tc"`
	Syn          bool   `json:"syn"`
	Fin          bool   `json:"fin"`
	Rst          bool   `json:"rst"`
	Psh          bool   `json:"psh"`
	Ack          bool   `json:"ack"`
	State        string `json:"state"`
	TsMaxRegions int    `json:"ts_max_regions"`
	TcMaxRegions int    `json:"tc_max_regions"`
}

type FlowRecord struct {
	Id            string           `json:"id,omitempty"`
	Start         types.CustomTime `json:"start"`
	End           types.CustomTime `json:"end"`
	SrcIp         string           `json:"src_ip"`
	DestIp        string           `json:"dest_ip"`
	DestPort      int              `json:"dest_port"`
	Proto         string           `json:"proto"`
	PktsToserver  int              `json:"pkts_toserver"`
	PktsToclient  int              `json:"pkts_toclient"`
	BytesToserver int              `json:"bytes_toserver"`
	BytesToclient int              `json:"bytes_toclient"`
}
