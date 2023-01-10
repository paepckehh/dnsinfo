# OVERVIEW

[paepche.de/dnsinfo](https://paepcke.de/dnsinfo)

-   Tired of remembering all the commandline switches of [dig|...] to [get|analyze] dns responses
-   Snapshot and compare dns anwers from different resolver for consistency
-   Get an summary overview report of all DNS types, even if your resolver blocks ANY
-   Get all kind of alerts about any anomalies, broken or depricated cryptographic functions
-   Support for all kind of custom or pre-defined custom and secure (DoT) (batch) resolver (see paepcke.de/dnsresolver)
-   100% pure go, minimal(internal-only) imports, use as app or api (see api.go), compatible with certinfo, tlsinfo, ...

# INSTALL
```
go install paepcke.de/dnsinfo/cmd/fdnsinfo@latest
```

### DOWNLOAD (prebuild)

[github.com/paepckehh/dnsinfo/releases](https://github.com/paepckehh/dnsinfo/releases)

# SHOWTIME

## Get a summary of a single host TLS [connection|handshake].

``` Shell
dnsinfo github.com
[Resolver:127.0.0.53:53] [DoT:false]
DNS A	github.com.	83861	IN	A	20.27.177.113
DNS CAA	github.com.	86209	IN	CAA	0 issue "digicert.com"
DNS CAA	github.com.	86209	IN	CAA	0 issue "globalsign.com"
DNS CAA	github.com.	86209	IN	CAA	0 issuewild "digicert.com"
DNS MX	github.com.	86222	IN	MX	10 alt4.aspmx.l.google.com.
DNS MX	github.com.	86222	IN	MX	5 alt1.aspmx.l.google.com.
DNS MX	github.com.	86222	IN	MX	5 alt2.aspmx.l.google.com.
DNS MX	github.com.	86222	IN	MX	1 aspmx.l.google.com.
DNS MX	github.com.	86222	IN	MX	10 alt3.aspmx.l.google.com.
DNS NS	github.com.	86198	IN	NS	dns4.p08.nsone.net.
DNS NS	github.com.	86198	IN	NS	ns-1283.awsdns-32.org.
DNS NS	github.com.	86198	IN	NS	ns-1707.awsdns-21.co.uk.
DNS NS	github.com.	86198	IN	NS	ns-421.awsdns-52.com.
DNS NS	github.com.	86198	IN	NS	ns-520.awsdns-01.net.
DNS NS	github.com.	86198	IN	NS	dns1.p08.nsone.net.
DNS NS	github.com.	86198	IN	NS	dns2.p08.nsone.net.
DNS NS	github.com.	86198	IN	NS	dns3.p08.nsone.net.
DNS SOA	github.com.	86211	IN	SOA	dns1.p08.nsone.net. hostmaster.nsone.net. 1656468023 43200 7200 1209600 3600
```

## Same but in ascii only non-color mode for post-processing, logging, ...

``` Shell
NO_COLOR=true dnsinfo github.com | grep ... 
[...]
```

## Need full details?

``` Shell
VERBOSE=true dnsinfo github.com 
[...]
```

# API

## Input Sources to Analyze:
## Output Format Styles via paepcke.de/reportstyle

-   Plain Text
-   Ansi Color Console
-   HTML
-   Custom \[get wild\]

## DNS state report of a single host.

``` Golang 
package main 

import ( 
	"os" 
	"paepcke.de/dnsinfo"
)

func main() { 
	os.Stdout.Write([]byte(HostReportAnsi("github.com"))) 
}

```

# DOCS

[pkg.go.dev/paepcke.de/dnsinfo](https://pkg.go.dev/paepcke.de/dnsinfo)

# CONTRIBUTION

Yes, Please! PRs Welcome! 
