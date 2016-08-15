package geo

import (
	"testing"

	"github.com/getlantern/goexpr"
	"github.com/stretchr/testify/assert"
)

func TestCITY_STATE(t *testing.T) {
	err := Init("geolite2-city.mmdb.gz")
	if !assert.NoError(t, err) {
		return
	}
	city := CITY_STATE(goexpr.Constant("66.69.229.169"))
	result := city.Eval(nil)
	assert.Equal(t, "Austin, Texas", result)
}
