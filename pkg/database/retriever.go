package database

import "github.com/clwg/eve-analyzer/pkg/model"

func (l *PostgresLogger) GetFlowLogsByDestIp(destIp string) ([]*model.FlowRecord, error) {
	rows, err := l.db.Query(
		`SELECT id, first_seen, last_seen, src_ip, dest_ip, dest_port, proto, bytes_to_server, bytes_to_client, pkts_to_server, pkts_to_client 
        FROM flow
        WHERE dest_ip = $1`,
		destIp,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []*model.FlowRecord
	for rows.Next() {
		var record model.FlowRecord
		err := rows.Scan(
			&record.Id,
			&record.Start,
			&record.End,
			&record.SrcIp,
			&record.DestIp,
			&record.DestPort,
			&record.Proto,
			&record.BytesToserver,
			&record.BytesToclient,
			&record.PktsToserver,
			&record.PktsToclient,
		)
		if err != nil {
			return nil, err
		}
		records = append(records, &record)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return records, nil
}

func (l *PostgresLogger) GetFlowLogsBySrcIp(srcIp string) ([]*model.FlowRecord, error) {
	rows, err := l.db.Query(
		`SELECT id, first_seen, last_seen, src_ip, dest_ip, dest_port, proto, bytes_to_server, bytes_to_client, pkts_to_server, pkts_to_client 
        FROM flow
        WHERE src_ip = $1`,
		srcIp,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []*model.FlowRecord
	for rows.Next() {
		var record model.FlowRecord
		err := rows.Scan(
			&record.Id,
			&record.Start,
			&record.End,
			&record.SrcIp,
			&record.DestIp,
			&record.DestPort,
			&record.Proto,
			&record.BytesToserver,
			&record.BytesToclient,
			&record.PktsToserver,
			&record.PktsToclient,
		)
		if err != nil {
			return nil, err
		}
		records = append(records, &record)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return records, nil
}

func (l *PostgresLogger) GetDNSQueryLogsByQname(qname string) ([]*model.DNSQuery, error) {
	rows, err := l.db.Query(
		`SELECT id, first_seen, last_seen, src_ip, dest_ip, dest_port, qname, count 
        FROM dnsquery
        WHERE qname LIKE $1`,
		qname,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []*model.DNSQuery
	for rows.Next() {
		var record model.DNSQuery
		err := rows.Scan(
			&record.ID,
			&record.FirstSeen,
			&record.LastSeen,
			&record.SrcIp,
			&record.DestIp,
			&record.DestPort,
			&record.Qname,
			&record.Count,
		)
		if err != nil {
			return nil, err
		}
		records = append(records, &record)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return records, nil
}

func (l *PostgresLogger) GetPassiveDNSLogsByQname(qname string) ([]*model.PassiveDNS, error) {
	rows, err := l.db.Query(
		`SELECT id, first_seen, last_seen, qname, rname, rtype, ttl, rdata, count 
        FROM passivedns
        WHERE qname LIKE $1`,
		qname,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []*model.PassiveDNS
	for rows.Next() {
		var record model.PassiveDNS
		err := rows.Scan(
			&record.ID,
			&record.FirstSeen,
			&record.LastSeen,
			&record.Qname,
			&record.RName,
			&record.RType,
			&record.TTL,
			&record.RData,
			&record.Count,
		)
		if err != nil {
			return nil, err
		}
		records = append(records, &record)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return records, nil
}

func (l *PostgresLogger) GetPassiveDNSLogsByRdata(rdata string) ([]*model.PassiveDNS, error) {
	rows, err := l.db.Query(
		`SELECT id, first_seen, last_seen, qname, rname, rtype, ttl, rdata, count 
        FROM passivedns
        WHERE rdata LIKE $1`,
		rdata,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []*model.PassiveDNS
	for rows.Next() {
		var record model.PassiveDNS
		err := rows.Scan(
			&record.ID,
			&record.FirstSeen,
			&record.LastSeen,
			&record.Qname,
			&record.RName,
			&record.RType,
			&record.TTL,
			&record.RData,
			&record.Count,
		)
		if err != nil {
			return nil, err
		}
		records = append(records, &record)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return records, nil
}
