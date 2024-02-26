package model

import (
	"github.com/clwg/eve-analyzer/types"
)

type Event struct {
	Timestamp types.CustomTime `json:"timestamp"`
	FlowID    int64            `json:"flow_id"`
	InIface   string           `json:"in_iface"`
	EventType string           `json:"event_type"`
	SrcIP     string           `json:"src_ip"`
	SrcPort   int              `json:"src_port"`
	DestIP    string           `json:"dest_ip"`
	DestPort  int              `json:"dest_port"`
	Proto     string           `json:"proto"`
	PktSrc    string           `json:"pkt_src"`
	DNS       DNS              `json:"dns,omitempty"`
	Flow      Flow             `json:"flow,omitempty"`
	TLS       TLS              `json:"tls,omitempty"`
	Quic      Quic             `json:"quic,omitempty"`
	TCP       TCP              `json:"tcp,omitempty"`
	HTTP      HTTP             `json:"http,omitempty"`
}
