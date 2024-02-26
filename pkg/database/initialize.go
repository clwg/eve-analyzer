package database

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var datasourceName string

var (
	instance *PostgresLogger
	once     sync.Once
)

type PostgresLogger struct {
	db *sql.DB
}

func SetDatasourceName(name string) {
	datasourceName = name
}

func PostgresEventLogger() *PostgresLogger {
	once.Do(func() {
		// Open a connection to the PostgreSQL database
		db, err := sql.Open("postgres", datasourceName)
		if err != nil {
			log.Fatalf("Could not open postgres database: %v", err)
		}

		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS flow (
			id uuid PRIMARY KEY,
			first_seen TIMESTAMP NOT NULL,
			last_seen TIMESTAMP NOT NULL,
			src_ip INET NOT NULL,
			dest_ip INET NOT NULL,
			dest_port INTEGER NOT NULL,
			proto TEXT NOT NULL,
			bytes_to_server BIGINT NOT NULL,
			bytes_to_client BIGINT NOT NULL,
			pkts_to_server BIGINT NOT NULL,
			pkts_to_client BIGINT NOT NULL
		);`)
		if err != nil {
			log.Fatalf("Could not create flow table: %v", err)
		}

		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS passivedns (
			id uuid PRIMARY KEY,
			first_seen TIMESTAMP NOT NULL,
			last_seen TIMESTAMP NOT NULL,
			qname TEXT NOT NULL,
			rname TEXT NOT NULL,
			rtype TEXT NOT NULL,
			ttl INTEGER NOT NULL,
			rdata TEXT NOT NULL,
			count BIGINT NOT NULL
		);`)
		if err != nil {
			log.Fatalf("Could not create PassiveDNS table: %v", err)
		}

		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS dnsquery (
			id uuid PRIMARY KEY,
			first_seen TIMESTAMP NOT NULL,
			last_seen TIMESTAMP NOT NULL,
			src_ip INET NOT NULL,
			dest_ip INET NOT NULL,
			dest_port INTEGER NOT NULL,
			qname TEXT NOT NULL,
			count BIGINT NOT NULL
		);`)
		if err != nil {
			log.Fatalf("Could not create PassiveDNS table: %v", err)
		}

		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tls (
			id uuid PRIMARY KEY,
			first_seen TIMESTAMP NOT NULL,
			last_seen TIMESTAMP NOT NULL,
			subject TEXT NOT NULL,
			issuerdn TEXT NOT NULL,
			serial TEXT NOT NULL,
			fingerprint TEXT NOT NULL,
			sni TEXT NOT NULL,
			version TEXT NOT NULL,
			notbefore TEXT NOT NULL,
			notafter TEXT NOT NULL,
			count BIGINT NOT NULL
		);`)
		if err != nil {
			log.Fatalf("Could not create tls table: %v", err)
		}

		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS sni_ip (
			id uuid PRIMARY KEY,
			first_seen TIMESTAMP NOT NULL,
			last_seen TIMESTAMP NOT NULL,
			sni TEXT NOT NULL,
			ip INET NOT NULL,
			count BIGINT NOT NULL
		);`)
		if err != nil {
			log.Fatalf("Could not create sni_ip table: %v", err)
		}

		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS quic (
			id uuid PRIMARY KEY,
			first_seen TIMESTAMP NOT NULL,
			last_seen TIMESTAMP NOT NULL,
			dest_ip INET NOT NULL,
			dest_port INTEGER NOT NULL,
			sni TEXT NOT NULL,
			ja3 TEXT NOT NULL,
			count BIGINT NOT NULL
		);`)
		if err != nil {
			log.Fatalf("Could not create quic table: %v", err)
		}

		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS quic (
			id uuid PRIMARY KEY,
			first_seen TIMESTAMP NOT NULL,
			last_seen TIMESTAMP NOT NULL,
			dest_ip INET NOT NULL,
			dest_port INTEGER NOT NULL,
			sni TEXT NOT NULL,
			ja3 TEXT NOT NULL,
			count BIGINT NOT NULL
		);`)
		if err != nil {
			log.Fatalf("Could not create quic table: %v", err)
		}

		instance = &PostgresLogger{db: db}
	})
	return instance
}
