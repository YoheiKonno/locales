package mzn_IR

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type mzn_IR struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	decimal                []byte
	group                  []byte
	minus                  []byte
	percent                []byte
	percentPrefix          []byte
	percentSuffix          []byte
	perMille               []byte
	timeSeparator          []byte
	inifinity              []byte
	currencies             [][]byte // idx = enum of currency code
	currencyPositivePrefix []byte
	currencyPositiveSuffix []byte
	currencyNegativePrefix []byte
	currencyNegativeSuffix []byte
	monthsAbbreviated      [][]byte
	monthsNarrow           [][]byte
	monthsShort            [][]byte
	monthsWide             [][]byte
	daysAbbreviated        [][]byte
	daysNarrow             [][]byte
	daysShort              [][]byte
	daysWide               [][]byte
	periodsAbbreviated     [][]byte
	periodsNarrow          [][]byte
	periodsShort           [][]byte
	periodsWide            [][]byte
	erasAbbreviated        [][]byte
	erasNarrow             [][]byte
	erasWide               [][]byte
	timezones              map[string][]byte
}

// New returns a new instance of translator for the 'mzn_IR' locale
func New() locales.Translator {
	return &mzn_IR{
		locale:            "mzn_IR",
		pluralsCardinal:   nil,
		pluralsOrdinal:    nil,
		decimal:           []byte{},
		group:             []byte{},
		minus:             []byte{},
		percent:           []byte{},
		perMille:          []byte{},
		timeSeparator:     []byte{0x3a},
		currencies:        [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		monthsAbbreviated: [][]uint8{[]uint8(nil), {0xda, 0x98, 0xd8, 0xa7, 0xd9, 0x86, 0xd9, 0x88, 0xdb, 0x8c, 0xd9, 0x87}, {0xd9, 0x81, 0xd9, 0x88, 0xd8, 0xb1, 0xdb, 0x8c, 0xd9, 0x87}, {0xd9, 0x85, 0xd8, 0xa7, 0xd8, 0xb1, 0xd8, 0xb3}, {0xd8, 0xa2, 0xd9, 0x88, 0xd8, 0xb1, 0xdb, 0x8c, 0xd9, 0x84}, {0xd9, 0x85, 0xd9, 0x87}, {0xda, 0x98, 0xd9, 0x88, 0xd8, 0xa6, 0xd9, 0x86}, {0xda, 0x98, 0xd9, 0x88, 0xd8, 0xa6, 0xdb, 0x8c, 0xd9, 0x87}, {0xd8, 0xa7, 0xd9, 0x88, 0xd8, 0xaa}, {0xd8, 0xb3, 0xd9, 0xbe, 0xd8, 0xaa, 0xd8, 0xa7, 0xd9, 0x85, 0xd8, 0xa8, 0xd8, 0xb1}, {0xd8, 0xa7, 0xda, 0xa9, 0xd8, 0xaa, 0xd8, 0xa8, 0xd8, 0xb1}, {0xd9, 0x86, 0xd9, 0x88, 0xd8, 0xa7, 0xd9, 0x85, 0xd8, 0xa8, 0xd8, 0xb1}, {0xd8, 0xaf, 0xd8, 0xb3, 0xd8, 0xa7, 0xd9, 0x85, 0xd8, 0xa8, 0xd8, 0xb1}},
		monthsWide:        [][]uint8{[]uint8(nil), {0xda, 0x98, 0xd8, 0xa7, 0xd9, 0x86, 0xd9, 0x88, 0xdb, 0x8c, 0xd9, 0x87}, {0xd9, 0x81, 0xd9, 0x88, 0xd8, 0xb1, 0xdb, 0x8c, 0xd9, 0x87}, {0xd9, 0x85, 0xd8, 0xa7, 0xd8, 0xb1, 0xd8, 0xb3}, {0xd8, 0xa2, 0xd9, 0x88, 0xd8, 0xb1, 0xdb, 0x8c, 0xd9, 0x84}, {0xd9, 0x85, 0xd9, 0x87}, {0xda, 0x98, 0xd9, 0x88, 0xd8, 0xa6, 0xd9, 0x86}, {0xda, 0x98, 0xd9, 0x88, 0xd8, 0xa6, 0xdb, 0x8c, 0xd9, 0x87}, {0xd8, 0xa7, 0xd9, 0x88, 0xd8, 0xaa}, {0xd8, 0xb3, 0xd9, 0xbe, 0xd8, 0xaa, 0xd8, 0xa7, 0xd9, 0x85, 0xd8, 0xa8, 0xd8, 0xb1}, {0xd8, 0xa7, 0xda, 0xa9, 0xd8, 0xaa, 0xd8, 0xa8, 0xd8, 0xb1}, {0xd9, 0x86, 0xd9, 0x88, 0xd8, 0xa7, 0xd9, 0x85, 0xd8, 0xa8, 0xd8, 0xb1}, {0xd8, 0xaf, 0xd8, 0xb3, 0xd8, 0xa7, 0xd9, 0x85, 0xd8, 0xa8, 0xd8, 0xb1}},
		erasAbbreviated:   [][]uint8{{0xd9, 0xbe, 0x2e, 0xd9, 0x85}, {0xd9, 0x85, 0x2e}},
		erasNarrow:        [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:          [][]uint8{{0xd9, 0x82, 0xd8, 0xa8, 0xd9, 0x84, 0x20, 0xd9, 0x85, 0xdb, 0x8c, 0xd9, 0x84, 0xd8, 0xa7, 0xd8, 0xaf}, {0xd8, 0xa8, 0xd8, 0xb9, 0xd8, 0xaf, 0x20, 0xd9, 0x85, 0xdb, 0x8c, 0xd9, 0x84, 0xd8, 0xa7, 0xd8, 0xaf}},
		timezones:         map[string][]uint8{"HKT": {0x48, 0x4b, 0x54}, "OEZ": {0x4f, 0x45, 0x5a}, "NZDT": {0x4e, 0x5a, 0x44, 0x54}, "CHAST": {0x43, 0x48, 0x41, 0x53, 0x54}, "TMT": {0x54, 0x4d, 0x54}, "WAT": {0x57, 0x41, 0x54}, "MESZ": {0x4d, 0x45, 0x53, 0x5a}, "HAST": {0x48, 0x41, 0x53, 0x54}, "SRT": {0x53, 0x52, 0x54}, "WEZ": {0x57, 0x45, 0x5a}, "ChST": {0x43, 0x68, 0x53, 0x54}, "AEDT": {0x41, 0x45, 0x44, 0x54}, "CST": {0x43, 0x53, 0x54}, "CHADT": {0x43, 0x48, 0x41, 0x44, 0x54}, "WITA": {0x57, 0x49, 0x54, 0x41}, "LHST": {0x4c, 0x48, 0x53, 0x54}, "ACWST": {0x41, 0x43, 0x57, 0x53, 0x54}, "COT": {0x43, 0x4f, 0x54}, "AEST": {0x41, 0x45, 0x53, 0x54}, "ADT": {0x41, 0x44, 0x54}, "SAST": {0x53, 0x41, 0x53, 0x54}, "WIT": {0x57, 0x49, 0x54}, "WIB": {0x57, 0x49, 0x42}, "GFT": {0x47, 0x46, 0x54}, "ARST": {0x41, 0x52, 0x53, 0x54}, "GMT": {0x47, 0x4d, 0x54}, "IST": {0x49, 0x53, 0x54}, "UYT": {0x55, 0x59, 0x54}, "NZST": {0x4e, 0x5a, 0x53, 0x54}, "WESZ": {0x57, 0x45, 0x53, 0x5a}, "WAST": {0x57, 0x41, 0x53, 0x54}, "PDT": {0x50, 0x44, 0x54}, "AKST": {0x41, 0x4b, 0x53, 0x54}, "WART": {0x57, 0x41, 0x52, 0x54}, "LHDT": {0x4c, 0x48, 0x44, 0x54}, "AKDT": {0x41, 0x4b, 0x44, 0x54}, "CDT": {0x43, 0x44, 0x54}, "AWST": {0x41, 0x57, 0x53, 0x54}, "JST": {0x4a, 0x53, 0x54}, "HNT": {0x48, 0x4e, 0x54}, "ACWDT": {0x41, 0x43, 0x57, 0x44, 0x54}, "CAT": {0x43, 0x41, 0x54}, "MYT": {0x4d, 0x59, 0x54}, "AWDT": {0x41, 0x57, 0x44, 0x54}, "BOT": {0x42, 0x4f, 0x54}, "VET": {0x56, 0x45, 0x54}, "HKST": {0x48, 0x4b, 0x53, 0x54}, "ACDT": {0x41, 0x43, 0x44, 0x54}, "EAT": {0x45, 0x41, 0x54}, "UYST": {0x55, 0x59, 0x53, 0x54}, "EDT": {0x45, 0x44, 0x54}, "OESZ": {0x4f, 0x45, 0x53, 0x5a}, "GYT": {0x47, 0x59, 0x54}, "BT": {0x42, 0x54}, "∅∅∅": {0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85}, "MEZ": {0x4d, 0x45, 0x5a}, "CLST": {0x43, 0x4c, 0x53, 0x54}, "ACST": {0x41, 0x43, 0x53, 0x54}, "AST": {0x41, 0x53, 0x54}, "MDT": {0x4d, 0x44, 0x54}, "ART": {0x41, 0x52, 0x54}, "JDT": {0x4a, 0x44, 0x54}, "ECT": {0x45, 0x43, 0x54}, "CLT": {0x43, 0x4c, 0x54}, "WARST": {0x57, 0x41, 0x52, 0x53, 0x54}, "HADT": {0x48, 0x41, 0x44, 0x54}, "PST": {0x50, 0x53, 0x54}, "HAT": {0x48, 0x41, 0x54}, "TMST": {0x54, 0x4d, 0x53, 0x54}, "EST": {0x45, 0x53, 0x54}, "COST": {0x43, 0x4f, 0x53, 0x54}, "SGT": {0x53, 0x47, 0x54}, "MST": {0x4d, 0x53, 0x54}},
	}
}

// Locale returns the current translators string locale
func (mzn *mzn_IR) Locale() string {
	return mzn.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'mzn_IR'
func (mzn *mzn_IR) PluralsCardinal() []locales.PluralRule {
	return mzn.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'mzn_IR'
func (mzn *mzn_IR) PluralsOrdinal() []locales.PluralRule {
	return mzn.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'mzn_IR'
func (mzn *mzn_IR) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'mzn_IR'
func (mzn *mzn_IR) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'mzn_IR'
func (mzn *mzn_IR) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'mzn_IR' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (mzn *mzn_IR) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	return []byte(s)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'mzn_IR' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (mzn *mzn_IR) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	return []byte(s)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'mzn_IR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (mzn *mzn_IR) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := mzn.currencies[currency]
	return append(append([]byte{}, symbol...), s...)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'mzn_IR'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (mzn *mzn_IR) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := mzn.currencies[currency]
	return append(append([]byte{}, symbol...), s...)
}

// FmtDateShort returns the short date representation of 't' for 'mzn_IR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (mzn *mzn_IR) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'mzn_IR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (mzn *mzn_IR) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'mzn_IR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (mzn *mzn_IR) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'mzn_IR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (mzn *mzn_IR) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'mzn_IR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (mzn *mzn_IR) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'mzn_IR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (mzn *mzn_IR) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'mzn_IR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (mzn *mzn_IR) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'mzn_IR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (mzn *mzn_IR) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtDateTimeShort returns the short date & time representation of 't' for 'mzn_IR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (mzn *mzn_IR) FmtDateTimeShort(t time.Time) []byte {

	b := mzn.FmtDateShort(t)
	b = append(b, ' ')
	b = append(b, mzn.FmtTimeShort(t)...)

	return b
}

// FmtDateTimeMedium returns the medium date & time representation of 't' for 'mzn_IR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (mzn *mzn_IR) FmtDateTimeMedium(t time.Time) []byte {

	b := mzn.FmtDateMedium(t)
	b = append(b, ' ')
	b = append(b, mzn.FmtTimeMedium(t)...)

	return b
}

// FmtDateTimeLong returns the long date & time representation of 't' for 'mzn_IR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (mzn *mzn_IR) FmtDateTimeLong(t time.Time) []byte {

	b := mzn.FmtDateLong(t)
	b = append(b, ' ')
	b = append(b, mzn.FmtTimeLong(t)...)

	return b
}

// FmtDateTimeFull returns the full date & time representation of 't' for 'mzn_IR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (mzn *mzn_IR) FmtDateTimeFull(t time.Time) []byte {

	b := mzn.FmtDateFull(t)
	b = append(b, ' ')
	b = append(b, mzn.FmtTimeFull(t)...)

	return b
}
