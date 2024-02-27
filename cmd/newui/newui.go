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
	"templates/flowlogs.html",
	"web/templates/dnsquerylogs.html",
	"templates/passivednslogs.html",
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

	data := struct {
		Qname      string
		DNSQueries interface{}
	}{
		Qname:      qname,
		DNSQueries: dnsQueries,
	}

	renderTemplate(w, "dnsquerylogs.html", data)
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

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8081",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe()
}
