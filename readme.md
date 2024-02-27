# eve-analyzer

A suricata eve.json log parser and web interface providing;

- per device data flow tracking.
- per device DNS queries metadata
- passive DNS database

Some key points:

- properly parses out the timestamp from the eve.json file
- uses the public suffix list to accurate extract the domain name from the DNS query

## Parser

```bash
go run eve_parser.go <path to eve.json>
 ```

The parser will automatically create the necessary database tables and should run without any additional configuration.

## WebUI Documentation

Currently a basic UI exists for querying the data.

![screenshot](https://github.com/clwg/eve-analyzer/blob/main/screenshot.png)

```bash
go run cmd/webui/webui.go
```
[http://localhost:8080/](http://localhost:8080/)

Wildcard searches using "*" are support for dns Qname, Rdata.