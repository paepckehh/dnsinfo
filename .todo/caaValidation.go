// package dnsinfo ...
package dnsinfo

//
// XXX  DO NOT USE [YET]
// RFC 8659 still lacks plausible / trusted mapping strategy and support for this kine of caa record usecase
//
// import
import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"strings"
)

// CAA TLS CONNECTION VALIDATION
//
// ValidateHost ...
func ValidateHost(host string, trustChains [][]*x509.Certificate) error {
	return GetLocalhostResolver(true).ValidateHost(host, trustChains)
}

// ValidateConn ...
func ValidateConn(host string, conn *tls.Conn) error {
	return GetLocalhostResolver(true).ValidateConn(host, conn)
}

// ValidateHost ...
func (r *Resolver) ValidateHost(host string, trustChains [][]*x509.Certificate) error {
	return ValidateChain(host, r.Lookup(host), trustChains)
}

// ValidateConn ...
func (r *Resolver) ValidateConn(host string, conn *tls.Conn) error {
	return ValidateChain(host, r.Lookup(host), conn.ConnectionState().VerifiedChains)
}

// ValidateChain ...
func ValidateChain(host string, caas []string, trustChains [][]*x509.Certificate) error {
	if len(trustChains) < 1 {
		return errors.New("[caa] Empty Trust Chain")
	}
	if len(caas) < 1 {
		return nil // there is no requirement to have dns caa record at all, do not fail here
	}
	var ok bool
	issuerMust, policyhostMust := false, false
	for _, caa := range caas {
		issuer, policyhost := "", ""
		policyhost, ok = strings.CutPrefix(caa, _policyhost)
		if ok {
			if policyhost == host {
				continue
			}
			policyhostMust = true
		}
		issuer, ok = strings.CutPrefix(caa, _issue)
		if !ok {
			continue
		}
		issuerMust = true
		for _, certs := range trustChains {
			for _, cert := range certs {
				if cert.IsCA {
					// RFC needs trused mapping policy update
					_ = issuer
					issuerMust = false
				}
			}
		}
	}
	if !issuerMust && !policyhostMust {
		return nil
	}
	reason := ""
	if issuerMust {
		reason += " issuer"
	}
	if policyhostMust {
		reason += " host"
	}
	return errors.New("[caa] policy violation:" + reason)
}
