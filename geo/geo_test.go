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

func TestREGION(t *testing.T) {
	e := REGION(goexpr.Constant("66.69.229.169"))
	result := e.Eval(nil)
	assert.Equal(t, "Texas", result)
}

func TestREGION_CITY(t *testing.T) {
	e := REGION_CITY(goexpr.Constant("66.69.229.169"))
	result := e.Eval(nil)
	assert.Equal(t, "Texas, Austin", result)
}

func TestCOUNTRY_CODE(t *testing.T) {
	e := COUNTRY_CODE(goexpr.Constant("66.69.229.169"))
	result := e.Eval(nil)
	assert.Equal(t, "US", result)
}
