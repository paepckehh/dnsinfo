// package dnsinfo
package dnsinfo

// import
import (
	"os"
)

// const
const (
	_whitespace         = ' '
	_tab                = '\t'
	_tabSep             = "\t"
	_lineFeed           = '\n'
	_false              = "false"
	_true               = "true"
	_space              = " "
	_dot                = "."
	_dotRune            = '.'
	_empty              = ""
	_issue              = "issue: "
	_errCode            = "unable to resolve: "
	_linefeed           = "\n"
	_sep                = "\t"
	_tcptls             = "tcp-tls"
	_tcp                = "tcp"
	_udp                = "udp"
	_six                = "6"
	_four               = "4"
	_dns                = "DNS "
	_policyhost         = "policy host:"
	_rfail              = "FAIL "
	_noAnswer           = "no answer"
	_errLookup          = " lookup failed : "
	_errUnsupportedType = "unsupported Type: "
	_errKeyPin          = "[dnsinfo] [tls keypin verification failed] "
	_errReverseAnswer   = "[dnsinfo] [reverse-lookup] invalid response from server "
	_errReverseLookup   = "[dnsinfo] [reverse-lookup] not a valid IP4 address: "
)

//
// LITTLE HELPER
//

// fqdn ...
func fqdn(name string) string {
	l := len(name)
	if l > 0 {
		last := name[l-1]
		if last == _dotRune {
			return name
		}
		return name + _dot
	}
	return "empty.hostname"
}

// isFile ...
func isFile(filename string) bool {
	fi, err := os.Lstat(filename)
	if err != nil {
		return false
	}
	return fi.Mode().IsRegular()
}

// isActiveBool ...
func isActiveBool(in bool) string {
	if in {
		return _true
	}
	return _false
}

// bracket ...
func bracket(in string) string {
	if len(in) > 0 {
		return "[" + in + "]"
	}
	return ""
}
