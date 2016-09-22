// Package isp provides isp lookup functions for IPv4 addresses
package isp

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync/atomic"

	"github.com/Workiva/go-datastructures/trie/yfast"
	"github.com/getlantern/goexpr"
)

var (
	entries = &atomic.Value{}
)

func init() {
	entries.Store(newTree())
}

func newTree() *yfast.YFastTrie {
	return yfast.New(uint32(0)) // 32 bit universe size for 32 bit IPv4
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

// Init initializes the isp package using a datafile from
// https://lite.ip2location.com/database/ip-asn.
func Init(datafile string) error {
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
			name:  row[4],
		}
		newEntries = append(newEntries, entry)
	}
	tree := newTree()
	tree.Insert(newEntries...)
	entries.Store(tree)

	return nil
}

// ISP returns the ISP name for a given IPv4 address
func ISP(ip goexpr.Expr) goexpr.Expr {
	return &isp{ip}
}

type isp struct {
	ip goexpr.Expr
}

func (e *isp) Eval(params goexpr.Params) interface{} {
	_ip := e.ip.Eval(params)
	switch ip := _ip.(type) {
	case string:
		ent, found := lookup(ip)
		if !found {
			return nil
		}
		return strings.TrimSpace(ent.name)
	}
	return nil
}

func (e *isp) String() string {
	return fmt.Sprintf("ISP(%v)", e.ip)
}

// ASN returns the ASN number for a given IPv4 address as an int
func ASN(ip goexpr.Expr) goexpr.Expr {
	return &asn{ip}
}

type asn struct {
	ip goexpr.Expr
}

func (e *asn) Eval(params goexpr.Params) interface{} {
	_ip := e.ip.Eval(params)
	switch ip := _ip.(type) {
	case string:
		ent, found := lookup(ip)
		if !found {
			return nil
		}
		return ent.asn
	}
	return nil
}

func (e *asn) String() string {
	return fmt.Sprintf("ASN(%v)", e.ip)
}

func lookup(ip string) (*entry, bool) {
	i := ipStringToInt(ip)
	if i == -1 {
		return nil, false
	}
	tree := entries.Load().(*yfast.YFastTrie)
	_e := tree.Predecessor(uint64(i))
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
