package model

import "github.com/clwg/eve-analyzer/types"

type DNS struct {
	Version     int         `json:"version"`
	Type        string      `json:"type"`
	ID          int         `json:"id"`
	Flags       string      `json:"flags"`
	QR          bool        `json:"qr"`
	RD          bool        `json:"rd"`
	RA          bool        `json:"ra"`
	Opcode      int         `json:"opcode"`
	RRName      string      `json:"rrname"`
	RRType      string      `json:"rrtype"`
	RCode       string      `json:"rcode"`
	Answers     []Answer    `json:"answers"`
	Authorities []Authority `json:"authorities"`
}

type Answer struct {
	RRName string `json:"rrname"`
	RRType string `json:"rrtype"`
	TTL    int    `json:"ttl"`
	RData  string `json:"rdata"`
}

type Authority struct {
	RRName string `json:"rrname"`
	RRType string `json:"rrtype"`
	TTL    int    `json:"ttl"`
	SOA    SOA    `json:"soa"`
}

type SOA struct {
	MName   string `json:"mname"`
	RName   string `json:"rname"`
	Serial  int    `json:"serial"`
	Refresh int    `json:"refresh"`
	Retry   int    `json:"retry"`
	Expire  int    `json:"expire"`
	Minimum int    `json:"minimum"`
}

type PassiveDNS struct {
	ID           string           `json:"id"`
	Timestamp    types.CustomTime `json:"timestamp"`
	FirstSeen    types.CustomTime `json:"first_seen,omitempty"`
	LastSeen     types.CustomTime `json:"last_seen,omitempty"`
	Qname        string           `json:"qname"`
	Domain       string           `json:"domain,omitempty"`
	DomainSuffix string           `json:"domain_suffix,omitempty"`
	RName        string           `json:"rname"`
	RType        string           `json:"rtype"`
	TTL          int              `json:"ttl"`
	RData        string           `json:"rdata"`
	Count        int              `json:"count,omitempty"`
}

type DNSQuery struct {
	ID           string           `json:"id"`
	Timestamp    types.CustomTime `json:"timestamp"`
	FirstSeen    types.CustomTime `json:"first_seen"`
	LastSeen     types.CustomTime `json:"last_seen"`
	SrcIp        string           `json:"src_ip"`
	DestIp       string           `json:"dest_ip"`
	DestPort     int              `json:"dest_port"`
	Qname        string           `json:"qname"`
	Domain       string           `json:"domain,omitempty"`
	DomainSuffix string           `json:"domain_suffix,omitempty"`
	Count        int              `json:"count,omitempty"`
}
