package handlers

import (
	"fmt"

	"github.com/clwg/eve-analyzer/pkg/model"
)

func HandleHttp(data model.Event) {
	fmt.Println("\nHTTP:")
	fmt.Println("Source:", data.SrcIP)
	fmt.Println("SourcePort:", data.SrcPort)
	fmt.Println("Destination:", data.DestIP)
	fmt.Println("DestinationPort:", data.DestPort)
	fmt.Println("Timestamp", data.Timestamp)
	fmt.Println("Hostname", data.HTTP.Hostname)
	fmt.Println("URL", data.HTTP.URL)
	fmt.Println("Method", data.HTTP.HTTPMethod)
	fmt.Println("UserAgent", data.HTTP.HTTPUserAgent)
	fmt.Println("StatusCode", data.HTTP.Status)
}
