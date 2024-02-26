package handlers

import (
	"fmt"

	"github.com/clwg/eve-analyzer/pkg/database"
	"github.com/clwg/eve-analyzer/pkg/model"
	"github.com/clwg/eve-analyzer/utils"
)

func HandleQuic(data model.Event) {
	dbLogger := database.PostgresEventLogger()

	if data.Quic.Ja3.Hash != "" || data.Quic.Sni != "" {
		compoundKey := data.SrcIP + data.DestIP + fmt.Sprint(data.DestPort) + data.Proto + data.Quic.Ja3.Hash + data.Quic.Sni

		uuid, err := utils.GenerateUUIDv5(compoundKey)
		if err != nil {
			fmt.Printf("Error generating UUID: %v\n", err)
			return
		}

		quicRecord := model.QuicRecord{
			Id:        uuid.String(),
			Timestamp: data.Timestamp,
			DestIp:    data.DestIP,
			DestPort:  data.DestPort,
			Sni:       data.Quic.Sni,
			Ja3:       data.Quic.Ja3.Hash,
		}

		if err := dbLogger.QuicLog(quicRecord); err != nil {
			panic(err)
		}

	}
}
