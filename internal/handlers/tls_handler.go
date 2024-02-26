package handlers

import (
	"fmt"
	"strings"

	"github.com/clwg/eve-analyzer/pkg/database"
	"github.com/clwg/eve-analyzer/pkg/model"
	"github.com/clwg/eve-analyzer/utils"
)

func HandleTLS(data model.Event) {
	dbLogger := database.PostgresEventLogger()

	compoundKey := data.TLS.Subject + data.TLS.Issuerdn + data.TLS.Serial + data.TLS.Fingerprint + data.TLS.Sni + data.TLS.Version + data.TLS.Notbefore + data.TLS.Notafter

	uuid, err := utils.GenerateUUIDv5(compoundKey)
	if err != nil {
		fmt.Printf("Error generating UUID: %v\n", err)
		return
	}
	data.TLS.Id = uuid.String()

	if err := dbLogger.TLSLog(data); err != nil {
		panic(err)
	}

	// SniIp mapping
	sniCompoundKey := data.TLS.Sni + data.SrcIP
	uuid, err = utils.GenerateUUIDv5(sniCompoundKey)
	if err != nil {
		fmt.Printf("Error generating UUID: %v\n", err)
		return
	}
	sniIp := model.SniIp{
		Id:        uuid.String(),
		Timestamp: data.Timestamp,
		Sni:       strings.ToLower(data.TLS.Sni),
		Ip:        data.DestIP,
	}

	if err := dbLogger.SniIpLog(sniIp); err != nil {
		panic(err)
	}

}
