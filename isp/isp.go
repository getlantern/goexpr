// Package isp provides isp lookup functions for IPv4 addresses
package isp

import (
	"fmt"
	"sync/atomic"

	"github.com/getlantern/goexpr"
)

var (
	provider atomic.Value
)

// Provider implements the actual looking up of ISP and ASN information.
type Provider interface {
	// ISP looks up the name of the Internet Service Provider corresponding to the
	// given ip.
	ISP(ip string) (isp string, found bool)

	// ASN looks up the Autonomous System Number corresponding to the given ip.
	ASN(ip string) (asn int, found bool)
}

// SetProvider sets the ISP data provider
func SetProvider(prov Provider) {
	provider.Store(prov)
}

func getProvider() Provider {
	return provider.Load().(Provider)
}

// ISP returns the ISP name for a given IPv4 address
func ISP(ip goexpr.Expr) goexpr.Expr {
	return &isp{ip}
}

type isp struct {
	ip goexpr.Expr
}

func (e *isp) Eval(params goexpr.Params) interface{} {
	_ip := e.ip.Eval(params)
	switch ip := _ip.(type) {
	case string:
		result, found := getProvider().ISP(ip)
		if !found {
			return nil
		}
		return result
	}
	return nil
}

func (e *isp) String() string {
	return fmt.Sprintf("ISP(%v)", e.ip)
}

// ASN returns the ASN number for a given IPv4 address as an int
func ASN(ip goexpr.Expr) goexpr.Expr {
	return &asn{ip}
}

type asn struct {
	ip goexpr.Expr
}

func (e *asn) Eval(params goexpr.Params) interface{} {
	_ip := e.ip.Eval(params)
	switch ip := _ip.(type) {
	case string:
		result, found := getProvider().ASN(ip)
		if !found {
			return nil
		}
		return result
	}
	return nil
}

func (e *asn) String() string {
	return fmt.Sprintf("ASN(%v)", e.ip)
}
