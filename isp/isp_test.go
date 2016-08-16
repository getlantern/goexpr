package isp

import (
	"testing"

	"github.com/getlantern/goexpr"
	"github.com/stretchr/testify/assert"
)

func init() {
	err := Init("ip2location_test_data.csv")
	if err != nil {
		panic(err)
	}
}

func TestConversion(t *testing.T) {
	assert.Equal(t, "66.69.229.169", ipIntToString(ipStringToInt("66.69.229.169")))
	assert.EqualValues(t, -1, ipStringToInt("13.342.23"))
}

func TestISPGoodIP(t *testing.T) {
	assert.Equal(t, "Time Warner Cable Internet LLC", ISP(goexpr.Constant("66.69.229.169")).Eval(nil))
}

func TestISPUnknownIP(t *testing.T) {
	assert.Nil(t, ISP(goexpr.Constant("104.131.61.10")).Eval(nil))
}

func TestISPInvalidIP(t *testing.T) {
	assert.Nil(t, ISP(goexpr.Constant("13.342.23")).Eval(nil))
}

func TestASNGoodIP(t *testing.T) {
	assert.Equal(t, 11427, ASN(goexpr.Constant("66.69.229.169")).Eval(nil))
}

func TestASNUnknownIP(t *testing.T) {
	assert.Nil(t, ASN(goexpr.Constant("104.131.61.10")).Eval(nil))
}

func TestASNInvalidIP(t *testing.T) {
	assert.Nil(t, ASN(goexpr.Constant("13.342.23")).Eval(nil))
}
