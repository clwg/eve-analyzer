package database

import "github.com/clwg/eve-analyzer/pkg/model"

func (l *PostgresLogger) PassiveDNSLog(message model.PassiveDNS) error {
	_, err := l.db.Exec(
		`INSERT INTO passivedns 
			(id, first_seen, last_seen, qname, domain, domain_suffix, rname, rtype, ttl, rdata, count) 
		VALUES 
			($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,1)
		ON CONFLICT (id) DO UPDATE SET 
			last_seen = $3,
			count = passivedns.count + 1;`,
		message.ID,
		message.Timestamp,
		message.Timestamp,
		message.Qname,
		message.Domain,
		message.DomainSuffix,
		message.RName,
		message.RType,
		message.TTL,
		message.RData,
	)
	return err
}

func (l *PostgresLogger) DNSQueryLog(message model.DNSQuery) error {
	_, err := l.db.Exec(
		`INSERT INTO dnsquery 
            (id, first_seen, last_seen, src_ip, dest_ip, dest_port, qname, count ) 
        VALUES 
            ($1,$2,$3,$4,$5,$6,$7,1)
        ON CONFLICT (id) DO UPDATE SET 
            first_seen = LEAST(dnsquery.first_seen, EXCLUDED.first_seen),
            last_seen = GREATEST(dnsquery.last_seen, EXCLUDED.last_seen),
            count = dnsquery.count + 1;`,
		message.ID,
		message.Timestamp,
		message.Timestamp,
		message.SrcIp,
		message.DestIp,
		message.DestPort,
		message.Qname,
	)
	return err
}

func (l *PostgresLogger) TLSLog(message model.Event) error {

	_, err := l.db.Exec(
		`INSERT INTO tls
            (id, first_seen, last_seen, subject, issuerdn, serial, fingerprint, sni, version, notbefore, notafter, count )
        VALUES
            ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,1)
        ON CONFLICT (id) DO UPDATE SET
            first_seen = LEAST(tls.first_seen, EXCLUDED.first_seen),
            last_seen = GREATEST(tls.last_seen, EXCLUDED.last_seen),
            count = tls.count + 1;`,
		message.TLS.Id,
		message.Timestamp,
		message.Timestamp,
		message.TLS.Subject,
		message.TLS.Issuerdn,
		message.TLS.Serial,
		message.TLS.Fingerprint,
		message.TLS.Sni,
		message.TLS.Version,
		message.TLS.Notbefore,
		message.TLS.Notafter,
	)
	return err
}

func (l *PostgresLogger) SniIpLog(message model.SniIp) error {

	_, err := l.db.Exec(
		`INSERT INTO sni_ip
            (id, first_seen, last_seen, sni, ip, count )
        VALUES
            ($1,$2,$3,$4,$5,1)
        ON CONFLICT (id) DO UPDATE SET
            first_seen = LEAST(sni_ip.first_seen, EXCLUDED.first_seen),
            last_seen = GREATEST(sni_ip.last_seen, EXCLUDED.last_seen),
            count = sni_ip.count + 1;`,
		message.Id,
		message.Timestamp,
		message.Timestamp,
		message.Sni,
		message.Ip,
	)
	return err
}

func (l *PostgresLogger) FlowLog(message model.FlowRecord) error {

	_, err := l.db.Exec(
		`INSERT INTO flow
            (id, first_seen, last_seen, src_ip, dest_ip, dest_port, proto, bytes_to_server, bytes_to_client, pkts_to_server, pkts_to_client )
        VALUES
            ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
        ON CONFLICT (id) DO UPDATE SET
            first_seen = LEAST(flow.first_seen, EXCLUDED.first_seen),
            last_seen = GREATEST(flow.last_seen, EXCLUDED.last_seen),
			bytes_to_server = flow.bytes_to_server + EXCLUDED.bytes_to_server,
			bytes_to_client = flow.bytes_to_client + EXCLUDED.bytes_to_client,
			pkts_to_server = flow.pkts_to_server + EXCLUDED.pkts_to_server,
			pkts_to_client = flow.pkts_to_client + EXCLUDED.pkts_to_client;`,
		message.Id,
		message.Start,
		message.End,
		message.SrcIp,
		message.DestIp,
		message.DestPort,
		message.Proto,
		message.BytesToserver,
		message.BytesToclient,
		message.PktsToserver,
		message.PktsToclient,
	)
	return err
}

func (l *PostgresLogger) QuicLog(message model.QuicRecord) error {

	_, err := l.db.Exec(
		`INSERT INTO quic
            (id, first_seen, last_seen, dest_ip, dest_port, sni, ja3, count )
        VALUES
            ($1,$2,$3,$4,$5,$6,$7,1)
        ON CONFLICT (id) DO UPDATE SET
            first_seen = LEAST(quic.first_seen, EXCLUDED.first_seen),
            last_seen = GREATEST(quic.last_seen, EXCLUDED.last_seen),
            count = quic.count + 1;`,
		message.Id,
		message.Timestamp,
		message.Timestamp,
		message.DestIp,
		message.DestPort,
		message.Sni,
		message.Ja3,
	)
	return err
}
