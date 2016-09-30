package maxmind

import (
	"fmt"
	"net"

	"github.com/getlantern/goexpr/isp"
	geoip2 "github.com/oschwald/geoip2-golang"
)

// NewProvider creates a new provider using the specified MaxMind GeoIP2 ISP
// datafile.
func NewProvider(datafile string) (isp.ISPProvider, error) {
	r, err := geoip2.Open(datafile)
	if err != nil {
		return nil, fmt.Errorf("Unable to open datafile at %v: %v", datafile, err)
	}
	prov := &provider{r}
	return prov, nil
}

type provider struct {
	r *geoip2.Reader
}

func (prov *provider) ISP(ip string) (string, bool) {
	isp, found := prov.lookup(ip)
	if !found {
		return "", false
	}
	return isp.Organization, isp.Organization != ""
}

func (prov *provider) ASN(ip string) (int, bool) {
	isp, found := prov.lookup(ip)
	if !found {
		return 0, false
	}
	return int(isp.AutonomousSystemNumber), isp.AutonomousSystemNumber != 0
}

func (prov *provider) lookup(ip string) (*geoip2.ISP, bool) {
	isp, err := prov.r.ISP(net.ParseIP(ip))
	if err != nil {
		return nil, false
	}
	return isp, isp != nil
}
