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

var templates = template.Must(template.ParseFiles(
	"web/templates/index.html",
	"web/templates/flowlogs.html",
	"web/templates/dnsquerylogs.html",
	"web/templates/passivednslogs.html",
))

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", nil)
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

	query := strings.Replace(qname, "%", "*", -1)

	data := struct {
		Query      string
		DNSQueries interface{}
	}{
		Query:      query,
		DNSQueries: dnsQueries,
	}

	renderTemplate(w, "dnsquerylogs.html", data)
}

func handlePassiveDNS(w http.ResponseWriter, r *http.Request) {
	qname := r.URL.Query().Get("qname")
	rdata := r.URL.Query().Get("rdata")

	db := database.PostgresEventLogger()
	var records interface{}
	var err error

	qname = strings.Replace(qname, "*", "%", -1)
	rdata = strings.Replace(rdata, "*", "%", -1)

	// Handle qname and rdata within the same route
	if qname != "" {
		records, err = db.GetPassiveDNSLogsByQname(qname)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if rdata != "" {
		records, err = db.GetPassiveDNSLogsByRdata(rdata)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	query := strings.Replace(qname, "%", "*", -1)

	data := struct {
		Query   string
		Records interface{}
	}{
		Query:   query,
		Records: records,
	}

	renderTemplate(w, "passivednslogs.html", data)
}

func handleFlow(w http.ResponseWriter, r *http.Request) {
	ip := r.URL.Query().Get("ip")
	destIP := r.URL.Query().Get("dest_ip")
	srcIP := r.URL.Query().Get("src_ip")

	db := database.PostgresEventLogger()
	var flowLogs interface{}
	var err error
	var query string

	// Handle dest_ip and src_ip within the same route
	if destIP != "" || ip != "" { // Prioritize dest_ip if both are provided
		flowLogs, err = db.GetFlowLogsByDestIp(destIP)
		query = destIP
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if srcIP != "" {
		flowLogs, err = db.GetFlowLogsBySrcIp(srcIP)
		query = srcIP
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	data := struct {
		Query    string
		FlowLogs interface{}
	}{
		Query:    query,
		FlowLogs: flowLogs,
	}

	renderTemplate(w, "flowlogs.html", data)
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

	// Serve static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))

	// Handle index
	r.HandleFunc("/", handleIndex)
	r.HandleFunc("/dnsquery", handleDNSQueryLogsByQname)
	r.HandleFunc("/passivedns", handlePassiveDNS)
	r.HandleFunc("/flow", handleFlow)

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8081",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe()
}
