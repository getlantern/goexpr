package geo

import (
	"testing"

	"github.com/getlantern/goexpr"
	"github.com/stretchr/testify/assert"
)

func init() {
	err := Init("geolite2-city.mmdb.gz")
	if err != nil {
		panic(err)
	}
}

func TestCITY(t *testing.T) {
	e := CITY(goexpr.Constant("66.69.229.169"))
	result := e.Eval(nil)
	assert.Equal(t, "Austin", result)
}

func TestSUBD(t *testing.T) {
	e := SUBD(goexpr.Constant("66.69.229.169"))
	result := e.Eval(nil)
	assert.Equal(t, "Texas", result)
}

func TestCITY_SUBD(t *testing.T) {
	e := CITY_SUBD(goexpr.Constant("66.69.229.169"))
	result := e.Eval(nil)
	assert.Equal(t, "Austin, Texas", result)
}

func TestCOUNTRY_CODE(t *testing.T) {
	e := COUNTRY_CODE(goexpr.Constant("66.69.229.169"))
	result := e.Eval(nil)
	assert.Equal(t, "US", result)
}
