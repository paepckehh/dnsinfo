// package dnsinfo ...
package dnsinfo

// import
import (
	"sort"
	"strings"

	"github.com/miekg/dns"
)

//
// REPORT
//

// report ...
func (r *Report) generateAll() string {
	var report strings.Builder
	report.WriteString(bracket("Resolver:" + r.Resolver.Server))
	report.WriteString(_space + bracket("DoT:"+isActiveBool(r.Resolver.DoT)))
	if r.Resolver.DoT {
		report.WriteString(_space + bracket("KeyPin:"+r.Resolver.TLSKeyPin))
	}
	report.WriteString(_linefeed)
	if err := r.Resolver.IsFunctional(); err != nil {
		report.WriteString("FAIL DNS Resolver Connect: " + err.Error() + _linefeed)
		return report.String()
	}
	answer, err := r.Resolver.Exchange(r.Query, r.Raw, r.Summary, r.Type)
	if err != nil {
		report.WriteString("FAIL DNS Resolver Exchange: " + err.Error() + _linefeed)
		return report.String()
	}
	if r.Summary {
		l := len(answer.Summary)
		sum := make([]string, l)
		for _, v := range answer.Summary {
			sum = append(sum, v)
		}
		sort.Strings(sum)
		report.Grow(l * 32)
		for _, v := range sum {
			report.WriteString(v)
		}
	}
	if r.Raw {
		for k, v := range answer.Raw {
			report.WriteString("\nDNS RAW TRACE DATA FOR [" + r.Query + "] [" + dns.TypeToString[k] + "]\n")
			report.WriteString(v)
		}
	}
	return report.String()
}
