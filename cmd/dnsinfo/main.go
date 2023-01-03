// package main
package main

// import
import (
	"io"
	"os"
	"syscall"

	"paepcke.de/dnsinfo"
	"paepcke.de/dnsresolver"
	"paepcke.de/reportstyle"
)

// const
const (
	_app      = "[dnsinfo] "
	_err      = "[error] "
	_html     = "HTML"
	_noColor  = "NO_COLOR"
	_verbose  = "VERBOSE"
	_forcedns = "FORCEDNS"
	_forcedot = "FORCEDOT"
	_linefeed = "\n"
)

// main ...
func main() {
	style := reportstyle.StyleAnsi()
	if isEnv(_noColor) {
		style = reportstyle.StyleText()
	}
	if isEnv(_html) {
		style = reportstyle.StyleHTML()
	}
	r := &dnsinfo.Report{
		Type:    dnsresolver.TypeAll,
		Summary: true,
		Style:   style,
	}
	if isEnv(_verbose) {
		r.Raw = true
	}
	switch {
	case isEnv(_forcedot):
		r.Resolver = dnsresolver.ResolverViaProvider(getEnv(_forcedot), true)
	case isEnv(_forcedns):
		r.Resolver = dnsresolver.ResolverViaProvider(getEnv(_forcedns), false)
	default:
		r.Resolver = dnsresolver.ResolverAuto()
	}
	switch {
	case isPipe():
		r.Query = getPipe()
		out(r.Generate())
	case isOsArgs():
		for i := 1; i < len(os.Args); i++ {
			r.Query = os.Args[i]
			out(r.Generate())
		}
	default:
		errExit("no pipe or input parameter found, example: dnsinfo ccc.de")
	}
}

//
// LITTLE GENERIC HELPER SECTION
//

// out ...
func out(msg string) {
	os.Stdout.Write([]byte(msg))
}

// errExit
func errExit(msg string) {
	out(_app + _err + msg + _linefeed)
	os.Exit(1)
}

// isPipe ...
func isPipe() bool {
	out, _ := os.Stdin.Stat()
	return out.Mode()&os.ModeCharDevice == 0
}

// getPipe ...
func getPipe() string {
	pipe, err := io.ReadAll(os.Stdin)
	if err != nil {
		errExit("reading data from pipe")
	}
	return string(pipe)
}

// isOsArgs ...
func isOsArgs() bool {
	return len(os.Args) > 1
}

// isEnv
func isEnv(in string) bool {
	if _, ok := syscall.Getenv(in); ok {
		return true
	}
	return false
}

// getEnv ...
func getEnv(in string) string {
	return os.Getenv(in)
}
