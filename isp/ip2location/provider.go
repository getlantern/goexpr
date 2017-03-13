package ip2location

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/Workiva/go-datastructures/trie/yfast"
	"github.com/getlantern/goexpr/isp"
)

// NewProvider creates a new provider using the specified datafile. This is a
// file from https://lite.ip2location.com/database/ip-asn.
func NewProvider(datafile string) (isp.Provider, error) {
	prov := &provider{
		yfast.New(uint32(0)), // 32 bit universe size for 32 bit IPv4
	}
	return prov, prov.init(datafile)
}

type entry struct {
	start uint64
	end   int64
	asn   int
	name  string
}

// Key implements the interface yfast.Entry
func (e *entry) Key() uint64 {
	return e.start
}

type provider struct {
	entries *yfast.YFastTrie
}

func (prov *provider) ISP(ip string) (string, bool) {
	ent, found := prov.lookup(ip)
	if !found {
		return "", false
	}
	return ent.name, true
}

func (prov *provider) ORG(ip string) (string, bool) {
	return prov.ISP(ip)
}

func (prov *provider) ASN(ip string) (int, bool) {
	ent, found := prov.lookup(ip)
	if !found {
		return 0, false
	}
	return ent.asn, true
}

func (prov *provider) ASName(ip string) (string, bool) {
	return prov.ISP(ip)
}

func (prov *provider) init(datafile string) error {
	file, err := os.Open(datafile)
	if err != nil {
		return fmt.Errorf("Unable to open datafile %v: %v", datafile, err)
	}
	defer file.Close()
	in := bufio.NewReader(file)
	csvIn := csv.NewReader(in)
	var newEntries []yfast.Entry
	for {
		row, err := csvIn.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("Unable to read data from %v: %v", datafile, err)
		}
		if len(row) != 5 {
			return fmt.Errorf("Unexpected row length %v: %v", row, err)
		}
		start, err := strconv.ParseInt(row[0], 10, 64)
		if err != nil {
			return fmt.Errorf("Unable to parse start IP %v: %v", row[0], err)
		}
		end, err := strconv.ParseInt(row[1], 10, 64)
		if err != nil {
			return fmt.Errorf("Unable to parse end IP %v: %v", row[1], err)
		}
		asn, err := strconv.Atoi(row[3])
		if err != nil {
			return fmt.Errorf("Unable to parse asn %v: %v", row[3], err)
		}
		entry := &entry{
			start: uint64(start),
			end:   end,
			asn:   asn,
			name:  strings.TrimSpace(row[4]),
		}
		newEntries = append(newEntries, entry)
	}
	prov.entries.Insert(newEntries...)
	return nil
}

func (prov *provider) lookup(ip string) (*entry, bool) {
	i := ipStringToInt(ip)
	if i == -1 {
		return nil, false
	}
	_e := prov.entries.Predecessor(uint64(i))
	if _e == nil {
		return nil, false
	}
	e, _ := _e.(*entry)
	if e.end < i {
		return nil, false
	}
	return e, true
}

// Below IP conversion functions taken from
// https://groups.google.com/forum/#!topic/golang-nuts/v4eJ5HK3stI

// Convert uint to string IP
func ipIntToString(ipnr int64) string {
	var bytes [4]byte
	bytes[0] = byte(ipnr & 0xFF)
	bytes[1] = byte((ipnr >> 8) & 0xFF)
	bytes[2] = byte((ipnr >> 16) & 0xFF)
	bytes[3] = byte((ipnr >> 24) & 0xFF)

	return fmt.Sprintf("%d.%d.%d.%d", bytes[3], bytes[2], bytes[1], bytes[0])
}

// Convert net.IP to int64
func ipStringToInt(ipnr string) int64 {
	bits := strings.Split(ipnr, ".")
	if len(bits) != 4 {
		return -1
	}

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum int64

	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)

	return sum
}
