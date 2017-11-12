# CEX - Consul EXplorer

Simple CLI tool for fast querying basic information from Consul.

Default CONSUL API endpoint is `http://localhost:8500`

To alter this value:
- export `CEX_CONSUL_URL` ENV variable with address of your consul api endpoint.
- use `-u` CLI flag

```
Usage:
  cex [command]

Available Commands:
  dc          Returns list of all discovarable datacenters
  help        Help about any command
  svc         Query services

Flags:
  -j, --json         JSON output
  -u, --url string   consul url (default "http://localhost:8500")
  -h, --help         help for cex

Use "cex [command] --help" for more information about a command.
```

```
Usage:
  cex svc [flags]

Flags:
  -d, --dc string        lookup specific datacenter
  -h, --help             help for svc
  -f, --hostname         return Node name instead of IP (in most cases that would be FQDN)
  -s, --service string   return list of hosts for requested service

Global Flags:
  -j, --json         JSON output
  -u, --url string   consul url (default "http://localhost:8500")
```

# Building from source

```
go get github.com/den-is/cex
cd $GOPATH/src/github.com/den-is/cex
go build
go install
```
After above steps copy `cex` executable from `$GOPATH/bin` into your `PATH`.
Or add `$GOPATH/bin` to your `PATH`.

# TODO
- Validate DC - if it is in list of discoverable DCs
- Validate service - if service exists in queried DC
- Include in output:
 - agent DC - DC you are connecting
 - queried DC - DC from which you are asking info about service. Defaults to agent DC.
- Use tags to refine service nodes output
