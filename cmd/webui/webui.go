package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/clwg/eve-analyzer/pkg/database"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var templates = template.Must(template.ParseFiles("templates/index.html", "templates/flowlogs.html", "templates/dnsquerylogs.html", "templates/passivednslogs.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", nil)
}

func handleFlow(w http.ResponseWriter, r *http.Request) {
	ip := r.URL.Query().Get("ip")
	destIP := r.URL.Query().Get("dest_ip")
	srcIP := r.URL.Query().Get("src_ip")

	db := database.PostgresEventLogger()
	var flowLogs interface{}
	var err error

	// Handle dest_ip and src_ip within the same route
	if destIP != "" || ip != "" { // Prioritize dest_ip if both are provided
		flowLogs, err = db.GetFlowLogsByDestIp(destIP)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if srcIP != "" {
		flowLogs, err = db.GetFlowLogsBySrcIp(srcIP)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	renderTemplate(w, "flowlogs.html", flowLogs)
}

func handleDNSQueryLogsByQname(w http.ResponseWriter, r *http.Request) {
	qname := r.URL.Query().Get("qname")
	qname = strings.Replace(qname, "*", "%", -1)

	db := database.PostgresEventLogger()
	dnsQueries, err := db.GetDNSQueryLogsByQname(qname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	renderTemplate(w, "dnsquerylogs.html", dnsQueries)
}

// Unified passivedns handler
func handlePassiveDNS(w http.ResponseWriter, r *http.Request) {
	qname := r.URL.Query().Get("qname")
	rdata := r.URL.Query().Get("rdata")

	db := database.PostgresEventLogger()
	var logs interface{}
	var err error

	qname = strings.Replace(qname, "*", "%", -1)
	rdata = strings.Replace(rdata, "*", "%", -1)

	// Handle qname and rdata within the same route
	if qname != "" {
		logs, err = db.GetPassiveDNSLogsByQname(qname)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if rdata != "" {
		logs, err = db.GetPassiveDNSLogsByRdata(rdata)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	renderTemplate(w, "passivednslogs.html", logs)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	database.SetDatasourceName(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	))

	r := mux.NewRouter()

	// Updated routes to use query parameters
	r.HandleFunc("/", handleIndex)
	r.HandleFunc("/flow", handleFlow)
	r.HandleFunc("/dnsquery", handleDNSQueryLogsByQname)
	r.HandleFunc("/passivedns", handlePassiveDNS)

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe()
}
