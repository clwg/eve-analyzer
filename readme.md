# eve-analyzer

A suricata eve.json log parser and web interface for querying flow and dns logs.

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
http://localhost:8080/

Wildcard searches using "*" are support for dns Qname, Rdata.


