package handlers

import (
	"fmt"
	"strings"

	"github.com/clwg/eve-analyzer/pkg/database"
	"github.com/clwg/eve-analyzer/pkg/domainsuffix"
	"github.com/clwg/eve-analyzer/pkg/model"
	"github.com/clwg/eve-analyzer/utils"
)

func HandleDNS(data model.Event) {

	dbLogger := database.PostgresEventLogger()

	if data.DNS.Type == "query" {

		compoundKey := data.SrcIP + data.DestIP + fmt.Sprint(data.DestPort) + data.DNS.RRName

		uuid, err := utils.GenerateUUIDv5(compoundKey)
		if err != nil {
			fmt.Printf("Error generating UUID: %v\n", err)
			return
		}

		var qname = strings.ToLower(data.DNS.RRName)

		domain, suffix, err := domainsuffix.ParseDomain(qname)
		if err != nil {
			fmt.Printf("Error parsing domain: %v\n", err)
			return
		}

		dnsQuery := model.DNSQuery{
			ID:           uuid.String(),
			Timestamp:    data.Timestamp,
			SrcIp:        data.SrcIP,
			DestIp:       data.DestIP,
			DestPort:     data.DestPort,
			Qname:        qname,
			Domain:       domain,
			DomainSuffix: suffix,
		}

		if err := dbLogger.DNSQueryLog(dnsQuery); err != nil {
			panic(err)
		}

	} else if data.DNS.Type == "answer" {

		for _, answer := range data.DNS.Answers {

			compoundKey := data.DNS.RRName + answer.RRName + answer.RRType + answer.RData

			uuid, err := utils.GenerateUUIDv5(compoundKey)
			if err != nil {
				fmt.Printf("Error generating UUID: %v\n", err)
				return
			}

			var qname = strings.ToLower(data.DNS.RRName)

			domain, suffix, err := domainsuffix.ParseDomain(qname)
			if err != nil {
				fmt.Printf("Error parsing domain: %v\n", err)
				return
			}

			passiveDNS := model.PassiveDNS{
				ID:           uuid.String(),
				Timestamp:    data.Timestamp,
				Qname:        qname,
				Domain:       domain,
				DomainSuffix: suffix,
				RName:        strings.ToLower(answer.RRName),
				RType:        answer.RRType,
				TTL:          answer.TTL,
				RData:        strings.ToLower(answer.RData),
			}

			if err := dbLogger.PassiveDNSLog(passiveDNS); err != nil {
				panic(err)
			}
		}
	}

}
