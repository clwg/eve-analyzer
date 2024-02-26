package consumer

import (
	"fmt"

	"github.com/clwg/eve-analyzer/internal/handlers"
	"github.com/clwg/eve-analyzer/pkg/model"
)

type HandlerFunc func(data model.Event)

func Dispatch(data model.Event) {
	handlers := map[string]HandlerFunc{
		"flow": handlers.HandleFlow,
		"dns":  handlers.HandleDNS,
		"tls":  handlers.HandleTLS,
		//"http": handlers.HandleHttp,
		"quic": handlers.HandleQuic,
	}

	if handler, ok := handlers[data.EventType]; ok {
		handler(data)
	} else {
		//fmt.Printf("No handler for type: %s\n", data.EventType)
	}
}

func ConsumeData(dataChan chan model.Event) {
	counter := 0
	for data := range dataChan {
		Dispatch(data)
		counter++
		if counter%100 == 0 {
			fmt.Printf("Processed events: %d\n", counter)
		}
	}
}
