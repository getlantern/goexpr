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

	// ORG looks up the name of the Organization corresponding to the given ip
	// (may be different than ISP).
	ORG(ip string) (org string, found bool)

	// ASN looks up the Autonomous System Number corresponding to the given ip.
	ASN(ip string) (asn int, found bool)
}

// SetProvider sets the ISP data provider
func SetProvider(prov Provider) {
	provider.Store(withCaching(prov, 1000000))
}

func getProvider() Provider {
	return provider.Load().(Provider)
}

// ISP returns the ISP name for a given IPv4 address
func ISP(ip goexpr.Expr) goexpr.Expr {
	return &ispExpr{"ISP", ip, func(ip string) (interface{}, bool) {
		return getProvider().ISP(ip)
	}}
}

// ORG returns the Organization name for a given IPv4 address (similar to ISP
// but may have different data depending on provider used).
func ORG(ip goexpr.Expr) goexpr.Expr {
	return &ispExpr{"ORG", ip, func(ip string) (interface{}, bool) {
		return getProvider().ORG(ip)
	}}
}

// ASN returns the ASN number for a given IPv4 address as an int
func ASN(ip goexpr.Expr) goexpr.Expr {
	return &ispExpr{"ASN", ip, func(ip string) (interface{}, bool) {
		return getProvider().ASN(ip)
	}}
}

type ispExpr struct {
	name string
	ip   goexpr.Expr
	fn   func(ip string) (interface{}, bool)
}

func (e *ispExpr) Eval(params goexpr.Params) interface{} {
	_ip := e.ip.Eval(params)
	switch ip := _ip.(type) {
	case string:
		result, found := e.fn(ip)
		if !found {
			return nil
		}
		return result
	}
	return nil
}

func (e *ispExpr) WalkParams(cb func(string)) {
	e.ip.WalkParams(cb)
}

func (e *ispExpr) WalkOneToOneParams(cb func(string)) {
	// this function is not one-to-one, stop
}

func (e *ispExpr) WalkLists(cb func(goexpr.List)) {
	e.ip.WalkLists(cb)
}

func (e *ispExpr) String() string {
	return fmt.Sprintf("%v(%v)", e.name, e.ip)
}
