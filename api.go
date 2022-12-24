// package dnsinfo provides dns reports
package dnsinfo

// import
import (
	"paepcke.de/dnsresolver"
	"paepcke.de/reportstyle"
)

//
// SIMPLE API
//

// ReportText ...
func ReportText(hostname string) string {
	r := &Report{
		Query:    hostname,
		Type:     dnsresolver.TypeAll,
		Resolver: dnsresolver.ResolverAuto(),
		Summary:  true,
		Style:    reportstyle.StyleText(),
	}
	return r.Generate()
}

// ReportAnsi ...
func ReportAnsi(hostname string) string {
	r := &Report{
		Query:    hostname,
		Type:     dnsresolver.TypeAll,
		Resolver: dnsresolver.ResolverAuto(),
		Summary:  true,
		Style:    reportstyle.StyleAnsi(),
	}
	return r.Generate()
}

// ReportMarkdown ...
func ReportMarkdown(hostname string) string {
	r := &Report{
		Query:    hostname,
		Type:     dnsresolver.TypeAll,
		Resolver: dnsresolver.ResolverAuto(),
		Summary:  true,
		Style:    reportstyle.StyleMarkdown(),
	}
	return r.Generate()
}

// ReportHTML ...
func ReportHTML(hostname string) string {
	r := &Report{
		Query:    hostname,
		Type:     dnsresolver.TypeAll,
		Resolver: dnsresolver.ResolverAuto(),
		Summary:  true,
		Style:    reportstyle.StyleHTML(),
	}
	return r.Generate()
}

//
//
// GENERIC BACKEND
//

// Report ...
type Report struct {
	// Hostname is the name to query
	Query string
	// Target is the DNS type(s) to query
	Type []uint16
	// Resolver is the DNS Nameserver profile for the query
	Resolver *dnsresolver.Resolver
	// Summary decides if the Summary is part of the Report
	Summary bool
	// Raw decides if the Raw Traces are part of the Report
	Raw bool
	// Style is the Report style (text, html, ...)
	Style *reportstyle.Style
}

// Generate ...
func (r *Report) Generate() string {
	return r.generateAll()
}
