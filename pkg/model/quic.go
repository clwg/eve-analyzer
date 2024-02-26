package model

import "github.com/clwg/eve-analyzer/types"

type Extension struct {
	Name   string   `json:"name,omitempty"`
	Type   int      `json:"type"`
	Values []string `json:"values,omitempty"`
}

type Ja3 struct {
	Hash   string `json:"hash"`
	String string `json:"string"`
}

type Quic struct {
	Version    string      `json:"version"`
	Sni        string      `json:"sni"`
	Ja3        Ja3         `json:"ja3"`
	Extensions []Extension `json:"extensions"`
}

type QuicRecord struct {
	Id        string           `json:"id"`
	Timestamp types.CustomTime `json:"timestamp"`
	DestIp    string           `json:"dest_ip"`
	DestPort  int              `json:"dest_port"`
	Sni       string           `json:"sni"`
	Ja3       string           `json:"ja3"`
}
