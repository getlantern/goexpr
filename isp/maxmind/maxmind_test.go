package maxmind

import (
	"testing"

	"github.com/getlantern/goexpr"
	"github.com/getlantern/goexpr/isp"
	"github.com/stretchr/testify/assert"
)

func init() {
	prov, err := NewProvider("GeoIP2-ISP.mmdb")
	if err != nil {
		panic(err)
	}
	isp.SetProvider(prov)
}

func TestISPGoodIP(t *testing.T) {
	assert.Equal(t, "Time Warner Cable", isp.ISP(goexpr.Constant("66.69.229.169")).Eval(nil))
}

func TestISPUnknownIP(t *testing.T) {
	assert.Nil(t, isp.ISP(goexpr.Constant("10.131.61.10")).Eval(nil))
}

func TestISPInvalidIP(t *testing.T) {
	assert.Nil(t, isp.ISP(goexpr.Constant("13.342.23")).Eval(nil))
}

func TestASNGoodIP(t *testing.T) {
	assert.Equal(t, 11427, isp.ASN(goexpr.Constant("66.69.229.169")).Eval(nil))
}

func TestASNUnknownIP(t *testing.T) {
	assert.Nil(t, isp.ASN(goexpr.Constant("10.131.61.10")).Eval(nil))
}

func TestASNInvalidIP(t *testing.T) {
	assert.Nil(t, isp.ASN(goexpr.Constant("13.342.23")).Eval(nil))
}
