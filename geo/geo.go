package geo

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/getlantern/goexpr"
	"github.com/getlantern/golog"
	geoip2 "github.com/oschwald/geoip2-golang"
)

const (
	dbURL = "https://geolite.maxmind.com/download/geoip/database/GeoLite2-City.mmdb.gz"
)

var (
	log = golog.LoggerFor("goexpr")

	db      *geoip2.Reader
	dbMutex sync.RWMutex
)

// Init initializes the Geolocation subsystem, storing the database file at the
// given dbFile location. It will periodically fetch updates from the maxmind
// website.
func Init(dbFile string) error {
	_db, dbDate, err := readDbFromFile(dbFile)
	if err != nil {
		_db, dbDate, err = readDbFromWeb(dbFile)
		if err != nil {
			return fmt.Errorf("Unable to read DB from file or web: %v", err)
		}
	}
	dbMutex.Lock()
	db = _db
	dbMutex.Unlock()
	go keepDbCurrent(dbFile, dbDate)
	return nil
}

func CITY_STATE(ip goexpr.Expr) goexpr.Expr {
	return &cityState{ip}
}

type cityState struct {
	ip goexpr.Expr
}

func (e *cityState) Eval(params goexpr.Params) interface{} {
	_ip := e.ip.Eval(params)
	switch ip := _ip.(type) {
	case string:
		city := cityFor(ip)
		if city == nil {
			return nil
		}
		cityName := city.City.Names["en"]
		if len(city.Subdivisions) > 0 {
			return cityName + ", " + city.Subdivisions[0].Names["en"]
		}
		return cityName
	}
	return nil
}

func (e *cityState) String() string {
	return fmt.Sprintf("ISP(%v)", e.ip)
}

func cityFor(ip string) *geoip2.City {
	parsedIP := net.ParseIP(ip)
	city, err := getDB().City(parsedIP)
	if err != nil {
		return nil
	}
	return city
}

func getDB() *geoip2.Reader {
	dbMutex.RLock()
	_db := db
	dbMutex.RUnlock()
	return _db
}

// readDbFromFile reads the MaxMind database and timestamp from a file
func readDbFromFile(dbFile string) (*geoip2.Reader, time.Time, error) {
	dbData, err := ioutil.ReadFile(dbFile)
	if err != nil {
		return nil, time.Time{}, fmt.Errorf("Unable to read db file %s: %s", dbFile, err)
	}
	fileInfo, err := os.Stat(dbFile)
	if err != nil {
		return nil, time.Time{}, fmt.Errorf("Unable to stat db file %s: %s", dbFile, err)
	}
	dbDate := fileInfo.ModTime()
	_db, err := openDb(dbData)
	if err != nil {
		return nil, time.Time{}, err
	}
	return _db, dbDate, nil
}

// keepDbCurrent checks for new versions of the database on the web every minute
// by issuing a HEAD request.  If a new database is found, this downloads it and
// submits it to server.dbUpdate for the run() routine to pick up.
func keepDbCurrent(dbFile string, dbDate time.Time) {
	for {
		time.Sleep(1 * time.Minute)
		headResp, err := http.Head(dbURL)
		if err != nil {
			log.Errorf("Unable to request modified of %s: %s", dbURL, err)
		}
		lm, err := lastModified(headResp)
		if err != nil {
			log.Errorf("Unable to parse modified date for %s: %s", dbURL, err)
		}
		if lm.After(dbDate) {
			log.Debug("Updating database from web")
			_db, _dbDate, err := readDbFromWeb(dbFile)
			if err != nil {
				log.Errorf("Unable to update database from web: %s", err)
			} else {
				dbDate = _dbDate
				dbMutex.Lock()
				db = _db
				dbMutex.Unlock()
			}
		}
	}
}

// readDbFromWeb reads the MaxMind database and timestamp from the web
func readDbFromWeb(dbFile string) (*geoip2.Reader, time.Time, error) {
	dbResp, err := http.Get(dbURL)
	if err != nil {
		return nil, time.Time{}, fmt.Errorf("Unable to get database from %s: %s", dbURL, err)
	}
	gzipDbData, err := gzip.NewReader(dbResp.Body)
	if err != nil {
		return nil, time.Time{}, fmt.Errorf("Unable to open gzip reader on response body%s", err)
	}
	defer gzipDbData.Close()
	dbData, err := ioutil.ReadAll(gzipDbData)
	if err != nil {
		return nil, time.Time{}, fmt.Errorf("Unable to fetch database from HTTP response: %s", err)
	}
	dbDate, err := lastModified(dbResp)
	if err != nil {
		return nil, time.Time{}, fmt.Errorf("Unable to parse Last-Modified header %s: %s", dbDate, err)
	}
	err = ioutil.WriteFile(dbFile, dbData, 0644)
	if err != nil {
		log.Errorf("Unable to save dbfile: %v", err)
	}
	db, err := openDb(dbData)
	if err != nil {
		return nil, time.Time{}, err
	}
	return db, dbDate, nil
}

// lastModified parses the Last-Modified header from a response
func lastModified(resp *http.Response) (time.Time, error) {
	lastModified := resp.Header.Get("Last-Modified")
	return http.ParseTime(lastModified)
}

// openDb opens a MaxMind in-memory db using the geoip2.Reader
func openDb(dbData []byte) (*geoip2.Reader, error) {
	db, err := geoip2.FromBytes(dbData)
	if err != nil {
		return nil, fmt.Errorf("Unable to open database: %s", err)
	}
	return db, nil
}
