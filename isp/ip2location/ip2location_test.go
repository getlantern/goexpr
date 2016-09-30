package ip2location

import (
	"testing"

	"github.com/getlantern/goexpr"
	"github.com/getlantern/goexpr/isp"
	"github.com/stretchr/testify/assert"
)

func init() {
	prov, err := NewProvider("ip2location_test_data.csv")
	if err != nil {
		panic(err)
	}
	isp.SetProvider(prov)
}

func TestConversion(t *testing.T) {
	assert.Equal(t, "66.69.229.169", ipIntToString(ipStringToInt("66.69.229.169")))
	assert.EqualValues(t, -1, ipStringToInt("13.342.23"))
}

func TestISPGoodIP(t *testing.T) {
	assert.Equal(t, "Time Warner Cable Internet LLC", isp.ISP(goexpr.Constant("66.69.229.169")).Eval(nil))
}

func TestISPUnknownIP(t *testing.T) {
	assert.Nil(t, isp.ISP(goexpr.Constant("104.131.61.10")).Eval(nil))
}

func TestISPInvalidIP(t *testing.T) {
	assert.Nil(t, isp.ISP(goexpr.Constant("13.342.23")).Eval(nil))
}

func TestORGGoodIP(t *testing.T) {
	assert.Equal(t, "Time Warner Cable Internet LLC", isp.ORG(goexpr.Constant("66.69.229.169")).Eval(nil))
}

func TestORGUnknownIP(t *testing.T) {
	assert.Nil(t, isp.ORG(goexpr.Constant("104.131.61.10")).Eval(nil))
}

func TestORGInvalidIP(t *testing.T) {
	assert.Nil(t, isp.ORG(goexpr.Constant("13.342.23")).Eval(nil))
}

func TestASNGoodIP(t *testing.T) {
	assert.Equal(t, 11427, isp.ASN(goexpr.Constant("66.69.229.169")).Eval(nil))
}

func TestASNUnknownIP(t *testing.T) {
	assert.Nil(t, isp.ASN(goexpr.Constant("104.131.61.10")).Eval(nil))
}

func TestASNInvalidIP(t *testing.T) {
	assert.Nil(t, isp.ASN(goexpr.Constant("13.342.23")).Eval(nil))
}
