package handlers

import (
	"fmt"

	"github.com/clwg/eve-analyzer/pkg/database"
	"github.com/clwg/eve-analyzer/pkg/model"
	"github.com/clwg/eve-analyzer/utils"
)

func HandleFlow(data model.Event) {
	dbLogger := database.PostgresEventLogger()

	compoundKey := data.SrcIP + data.DestIP + fmt.Sprint(data.DestPort) + data.Proto
	uuid, err := utils.GenerateUUIDv5(compoundKey)
	if err != nil {
		fmt.Printf("Error generating UUID: %v\n", err)
		return
	}

	flowRecord := model.FlowRecord{
		Id:            uuid.String(),
		SrcIp:         data.SrcIP,
		DestIp:        data.DestIP,
		DestPort:      data.DestPort,
		Proto:         data.Proto,
		Start:         data.Flow.Start,
		End:           data.Flow.End,
		BytesToserver: data.Flow.BytesToserver,
		BytesToclient: data.Flow.BytesToclient,
		PktsToserver:  data.Flow.PktsToserver,
		PktsToclient:  data.Flow.PktsToclient,
	}
	if err := dbLogger.FlowLog(flowRecord); err != nil {
		panic(err)
	}

}
