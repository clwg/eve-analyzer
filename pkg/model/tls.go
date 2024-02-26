package model

import "github.com/clwg/eve-analyzer/types"

type TLS struct {
	Id          string `json:"id,omitempty"`
	Subject     string `json:"subject"`
	Issuerdn    string `json:"issuerdn"`
	Serial      string `json:"serial"`
	Fingerprint string `json:"fingerprint"`
	Sni         string `json:"sni"`
	Version     string `json:"version"`
	Notbefore   string `json:"notbefore"`
	Notafter    string `json:"notafter"`
}

type SniIp struct {
	Id        string           `json:"id,omitempty"`
	Timestamp types.CustomTime `json:"timestamp"`
	Sni       string           `json:"sni"`
	Ip        string           `json:"ip"`
}
